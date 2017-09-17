// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cancel_test

import (
	"context"
	"time"

	"github.com/GoLangsam/do/cli/cancel"
)

func mainOperation(ctx context.Context) {
	time.Sleep(33 * time.Millisecond) // be busy
	println("Working completed")      // celebrate
	return                            // go home
}

func doneMainOperation(ctx context.Context) <-chan struct{} {
	cha := make(chan struct{})
	go func() {
		defer close(cha)
		time.Sleep(33 * time.Millisecond) // be busy
		println("Working completed")      // celebrate
		cha <- struct{}{}                 // signal 'done' (& go home)
	}()
	return cha
}

func ExampleReStart() {
	cancel.ReStart() // no need in Your main.go
}

func ExampleBackGround() {
	cancel.ReStart() // no need in Your main.go

	root := cancel.BackGround()
	ctx, cancel := context.WithCancel(root) // use it to derive Your own child context

	mainOperation(ctx)

	<-ctx.Done()
	cancel()
}

func ExampleWithCancel() {
	cancel.ReStart() // no need in Your main.go

	ctx, cancel := cancel.WithCancel()

	mainOperation(ctx) // pretend to be busy

	cancel()
	<-ctx.Done()
}

func ExampleWithTimeout() {
	cancel.ReStart() // no need in Your main.go

	ctx, cancel := cancel.WithTimeout(100 * time.Millisecond)

	mainOperation(ctx) // pretend to be busy

	cancel() // mainOperation finished: no more need to await timeout
	<-ctx.Done()
}

func ExampleWithDeadline() {
	cancel.ReStart() // no need in Your main.go

	ctx, _ := cancel.WithDeadline(time.Now().Add(100 * time.Millisecond))

	done := doneMainOperation(ctx) // pretend to be busy 'till dawn

	select {
	case <-done:
		println("Mission completed")
	case <-ctx.Done():
		println("Mission aborted")
	}
}
