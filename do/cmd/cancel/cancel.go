package cancel

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

//!+1
var cancel = make(chan struct{})
var kill = make(chan os.Signal)
var interrupt = make(chan os.Signal)
var ctrlC = make(chan os.Signal)

func init() {
	// Cancel traversal when input is detected.
	go func() {
		defer close(cancel)
		os.Stdin.Read(make([]byte, 1)) // read a single byte
	}()
	signal.Notify(kill, os.Kill)           // TODO: support
	signal.Notify(interrupt, os.Interrupt) // TODO: support
	signal.Notify(ctrlC, syscall.SIGHUP)
}

// usage:
//	select {
//	case <-cancel: // abort (after drain, if need)
//	case <-ctrl_C: // abort (after drain, if need)
//

// exit aborts the program with a hard exit(1)
func exit() {
	println("Program aborted!")
	os.Exit(1)
}

// Canceler launches a go routine which checks every n milliSeconds (n > 10)
// and terminates via os.Exit(1) if SIGHUP (Ctrl-C) or Enter are received.
func Canceler(ns ...int) {
	n := 100
	if len(ns) > 0 {
		n = ns[0]
	}
	if n < 11 {
		n = 11
	}
	go func() {
		for {
			select {
			case <-cancel:
				exit()
			case <-ctrlC:
				exit()
			default:
				time.Sleep(time.Duration(n) * time.Millisecond)
			}
		}
	}()
}

// Cancelled is a convenient alternative to brute force Canceler
func Cancelled() bool {
	select {
	case <-cancel:
		return true
	case <-ctrlC:
		return true
	default:
		return false
	}
}
