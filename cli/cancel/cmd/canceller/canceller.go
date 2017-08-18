// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package main is a toy to play with do/cli/cancel (and context)
package main

import (
	"fmt"
	"time"

	"github.com/golangsam/do/cli/cancel"
)

func main() {
	fmt.Println("Waiting for cancel signal:")
	fmt.Println("Press 'Enter' or 'Ctrl-C'!")
	timeout := 5 * time.Second
	fmt.Println("You have " + timeout.String() + " to go ...")

	// beg of use
	ctx, _ := cancel.WithTimeout(timeout)
	<-ctx.Done()
	// end of use

	fmt.Println("I see: " + ctx.Err().Error())
}
