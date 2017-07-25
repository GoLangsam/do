// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cancel provides cancellation primitives for cli based cmd's.
package cancel

import (
	"context"
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

// cancelled is a convenience
func cancelled() bool {
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

// exit aborts the program with a hard exit(1) after printing the error
func exit(err error) {
	println(err.Error())
	os.Exit(1)
}

// Canceler launches a go routine which checks every n milliSeconds (n > 10)
// and terminates via os.Exit(1) if Cancelled() returns true.
func Canceler(ns ...int) (ctx context.Context) {
	ctx, cancel := context.WithCancel(context.Background())

	n := 100
	if len(ns) > 0 {
		n = ns[0]
	}
	if n < 11 {
		n = 11
	}
	go func(ns int, cancel context.CancelFunc) {
		for {
			if cancelled() {
				cancel()
				return
			} else {
				time.Sleep(time.Duration(ns) * time.Millisecond)
			}
		}
	}(n, cancel)

	go func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				exit(ctx.Err())
			}
		}
	}(ctx)

	return ctx
}
