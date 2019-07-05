// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cancel - convenient cancellation for cli based cmd's.
//
// Function names match the underlying context package -
// just: their call signatures do not need any parent context.
//
// And: they build upon each other: You'll need only one.
//
// Choose exactly one from:
//  - WithCancel()      - includes BackGround, and listen
//  - WithTimeout(...)  - includes WithCancel, and waiter
//  - WithDeadline(...) - includes WithCancel, and waiter
// to get Your observed base context.
//
// When You like to switch-off any cancellation (temporarily),
// You may use
//  - BackGround()
// as a convenience (without any need to adjust Your import statements).
//
// Note: "Don't `panic`"!
// Extra checks protect You against accidental misuse
// and `panic` upon any subsequent use,
// unless You explicitly call `ReStart()` before.
//
// Technical note: WithCancel, WithTimeout and WithDeadline return to You
//  - a context (child =Timeout/Deadline) - You may derive further from it
//  - a CancelFunc (parent = Cancel) in case You dare to use it.
// Just: Most often You may safely ignore the CancelFunc,
// as `listen` and `waiter` take care of it.
// Thus, You need not to worry about any leaking.
//
// Be a happy gopher :-)
package cancel

import (
	"context"
	"os"
	"os/signal"
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

// ReStart permits You to call any of the `WithXyz`-functions again.
func ReStart() {
	haveBeenCalled = false
}

// event channels
var (
	newline = make(chan struct{})
	signals = make(chan os.Signal)
)

func init() {

	go func() { // spawn input listener
		defer close(newline)           // close cancel when input (via `Enter` pressed) is detected
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		newline <- struct{}{}          // signal cancel
	}()
	signal.Notify(signals, os.Kill, os.Interrupt)
}

// listen applies the given CancelFunc iff some cancellation is seen by cancelled()
func listen(cancel context.CancelFunc, msDelay ...int) {
	defer cancel()
	for {
		select {
		case <-newline:
			return
		case <-signals:
			return
		default:
			time.Sleep(100 * time.Millisecond) // as in
		}
	}
}

// waiter applies the returned parent-CancelFunc, iff
//  - internal parent WithCancel-context or
//  - returned WithDeadline-child context
// become canceled
func waiter(one, two context.Context, cancel context.CancelFunc) {
	defer cancel()
	select {
	case <-one.Done():
		return
	case <-two.Done():
		return
	}
}

// BackGround just gives You an initial root context (context.Background())
// without any functionality.
//
// Note: only exported as convenience - similar to the context package.
// Use one of the more advanced `WithXyz`-functions instead.
func BackGround() (ctx context.Context) {
	mustBeFirstCall()
	return context.Background()
}

// WithCancel gives You an initial root context (and it's CancelFunc)
// and spawns `listen` which cancels the context upon any cancellation event.
func WithCancel() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(BackGround())

	go listen(cancel)

	return ctx, cancel
}

// WithDeadline gives You an initial root context (and a CancelFunc)
// and spawns `waiter` which cancels the context upon
//  - any cancellation event
//  - deadline expires
// whichever is seen first
func WithDeadline(deadline time.Time) (ctx context.Context, cancel context.CancelFunc) {
	parent, cancel := WithCancel()

	var noNeed context.CancelFunc
	ctx, noNeed = context.WithDeadline(parent, deadline) // CancelFunc - no need
	_ = noNeed                                           // silence go vet

	go waiter(parent, ctx, cancel)

	return ctx, cancel
}

// WithTimeout gives You an initial rooted context (and a CancelFunc)
// and spawns `waiter` which cancels the context upon
//  - any cancellation event
//  - timeout elapses
// whichever is seen first
// Note: WithTimeout is simply a convenience for `WithDeadline(time.Now().Add(timeout))`.
func WithTimeout(timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	return WithDeadline(time.Now().Add(timeout))
}

// Done returns a do.Nok which returns false
// unless the given context has been cancelled.
//
// Done is a convenience.
func Done(ctx context.Context) func() bool {
	return func() bool {
		select {
		case <-ctx.Done():
			return true
		default:
		}
		return false
	}
}
