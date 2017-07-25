// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cancel provides cancellation primitives for cli based cmd's.
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

	go func() { // spawn input listener
		defer close(cancel)            // close cancel when input (via `Enter` pressed) is detected
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cancel <- struct{}{}           // signal cancel
	}()
	signal.Notify(kill, os.Kill)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(ctrlC, syscall.SIGHUP)
}

// exit aborts the program with a hard exit(1)
func exit() {
	println("Program aborted!")
	os.Exit(1)
}

// Canceler launches a go routine which checks every n milliSeconds (n > 10)
// and terminates via os.Exit(1) if Cancelled() returns true.
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
			if Cancelled() {
				exit()
			} else {
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
	case <-kill:
		return true
	case <-interrupt:
		return true
	default:
		return false
	}
}
