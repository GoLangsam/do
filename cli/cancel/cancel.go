// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cancel provides convenient cancellation primitives for cli based cmd's.
//
// The exposed names are inspired by the underlying context package -
// just: signatures differ slightly.
// And: they build upon each other, so You'll need only one.
//
// Choose exactly one from:
// - WithCancel   (includes BackGround, and listen)
// - WithTimeout  (includes WithCancel, and waiter)
// - WithDeadline (includes WithCancel, and waiter)
// to get Your observed root context.
//
// Or, if You like to switch-off any cancellation functionality (temporarily),
// You may use
// - BackGround
// as a convenience (without any need to adjust Your import statements).
//
// Note: "Don't `panic`"!
// Extra checks protect You against accidental misuse
// and `panic` upon any subsequent use,
// unless You explicitly call `ReStart()` before.
//
// Technical note: WithCancel, WithTimeout and WithDeadline return to You
// - their (child =Timeout/Deadline) context, so You may derive from it further
// - their parents (= Cancel) CancelFunc (in case You dare to use it).
// Just: Most often You may safely ignore the CancelFunc,
// as listen and waiter take care of it.
// Thus, You need not to worry about any leaking.
// Be a happy gopher :-)
package cancel

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Protect against accidental misuse - "Don't `panic`" :-)

var haveBeenCalled bool

func mustBeFirstCall() {
	if haveBeenCalled {
		panic("Only one call of some WithXyz! Hint: use ReStart(), if You really need me again!")
	}
	haveBeenCalled = true
}

// ReStart permits You to call any of the WithXyz funtion again.
func ReStart() {
	haveBeenCalled = false
}

// event channels
var (
	cancel    = make(chan struct{})  // Enter
	kill      = make(chan os.Signal) // os.Kill
	interrupt = make(chan os.Signal) // os.Interrupt
	ctrlC     = make(chan os.Signal) // syscall.SIGHUP
)

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

// listen applies the given CancelFunc iff some cancellation is seen by cancelled()
func listen(cancel context.CancelFunc, msDelay ...int) {
	defer cancel()
	for {
		if cancelled() {
			return
		} else {
			time.Sleep(100 * time.Millisecond) // as in
		}
	}
}

// waiter applies the returned parent-CancelFunc, iff
// internal parent WithCancel-context or
// returned WithDeadline-child context
// become canceled
func waiter(deadline time.Time) (ctx context.Context, cancel context.CancelFunc) {
	parent, cancel := WithCancel()
	ctx, _ = context.WithDeadline(parent, deadline) // here we need "context"!
	// But no dedicated CancelFunc

	go func(one, two context.Context, cancel context.CancelFunc) {
		defer cancel()
		select {
		case <-one.Done():
			return
		case <-two.Done():
			return
		}
	}(parent, ctx, cancel)

	return ctx, cancel
}

// BackGround just gives You an initial root context (context.Background())
// without any further functionalities.
//
// Note: BackGround is only exported as convenient similarity with the context package.
// You are encouraged to use one of the more advanced `WithXyz`-functions instead.
func BackGround() (ctx context.Context) {
	mustBeFirstCall()
	return context.Background()
}

// WithCancel gives You an initial root context (and it's CancelFunc)
// and closes it automatically upon any cancellation event.
//
// Thus, simply have Your main operation invoke the returned CancelFunc upon completition,
//   (Hint: You may like to defer the returned CancelFunc there, but You don't have to).
// and simply await `<-ctx.Done()` (*un*conditionally!) just before Your `main()` exits.
//
// Or You just `select` to wait for either:
// - a `<-ctx.Done()` in order to shutdown gracefully
// - a signal of sucessfull termination from your main operation
func WithCancel() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(BackGround())
	go listen(cancel)
	return ctx, cancel
}

// WithDeadline gives You an initial root context (and a CancelFunc, and an error channel)
// and closes it automatically upon
// - any cancellation event
// - deadline expires
// whichever is seen first
func WithDeadline(deadline time.Time) (ctx context.Context, cancel context.CancelFunc) {
	return waiter(deadline)
}

// WithTimeout gives You an initial rooted context (and a CancelFunc, and an error channel)
// and closes it automatically upon
// - any cancellation event
// - timeout elapses
// whichever is seen first
// Note: WithTimeout is simply a convenience for `WithDeadline(time.Now().Add(timeout))`.
func WithTimeout(timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	return WithDeadline(time.Now().Add(timeout))
}
