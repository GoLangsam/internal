package snip

import (
	"path/filepath"
	"io/ioutil"
	"os"
)

// Nice pattern for File-Writer - from "github.com\derekchiang\Sieve-of-Eratosthenes\soe.go"
// with noisy Close

func sampleWrite() {
	output, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	// write sth into output
}

// Following sample uses the "comma-error" style instead of panicing
//
// writeFile writes the given contents to the given path, creating any necessary parent directories.
// This is useful because both problem files and solution files may have directory structures.
func writeFile(path, contents string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(contents), 0644)
}
