// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 251.

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

// The du4 command computes the disk usage of the files in a directory.
package main

// The du4 variant includes cancellation:
// it terminates quickly when the user hits return.

import (
	"flag"

	"fmt"
//	"io/ioutil"

	"os"
	"path/filepath"
	"strings"
	"sort"
	"sync"
	"time"
)

type DirName struct{
	name string
	root string
}

func printNameUsage( dirnames map[string][]string ) {
	var names []string
	for name := range dirnames {
		names = append( names, name )
	}
	sort.Strings( names )
	for _, name := range names {
		if !*dFlag || len( dirnames[name] ) > 1 {
			fmt.Printf( "%v\t%v\n", name, strings.Join( dirnames[name], "\t") )
		}
	}

}


var vFlag = flag.Bool("v", false, "show verbose progress messages")
var dFlag = flag.Bool("d", false, "show only duplicate/multiple names")

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
func checkdone() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
}


//!-1

func main() {
	flag.Parse()			// ...determine roots...
	roots := flag.Args()		// Determine the initial directories.
	if len(roots) == 0 { roots = []string{"."} }

	go checkdone()			// Cancel traversal if input is detected.

	DirNames := make(map[string][]string )

	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan DirName)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, ".", &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(10 * time.Millisecond)
	}

loop:
	for {
		select {
		case <-done:
			for range fileSizes { }	// Drain fileSizes to allow existing goroutines to finish.
			return
		case dirname, ok := <-fileSizes:
			if !ok { break loop }	// fileSizes was closed
			DirNames[dirname.name] = append(DirNames[dirname.name], dirname.root)
		case <-tick:
			fmt.Print( "." )
		}
	}
	fmt.Println( "" )

	printNameUsage( DirNames ) // final totals
//	printDiskUsage(nfiles, nbytes) // final totals
	//!+
	// ...select loop...
}


func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(root, dir string, n *sync.WaitGroup, fileSizes chan<- DirName) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(filepath.Join(root, dir)) {
		if entry.IsDir() {
			var d = new(DirName)
			d.root = dir
			d.name = entry.Name()
			fileSizes <- *d
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, fileSizes)
		}
	}
}

var sema = make(chan struct{}, 20) // concurrency-limiting counting semaphore

// dirents returns the entries of directory dir.
//!+5
func dirents(dir string) []os.FileInfo {
	select {
		case sema <- struct{}{}: // acquire token
		case <-done:
			return nil // cancelled
	}
	defer func() { <-sema }() // release token

	// ...read directory...
	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}
