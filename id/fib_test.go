// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id_test

import (
	"fmt"

	"github.com/GoLangsam/do/id"
)

func Example_Fib() {
	for fib := range id.Fib(20) {
		fmt.Println(fib)
	}
	// Output:
	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// 55
	// 89
	// 144
	// 233
	// 377
	// 610
	// 987
	// 1597
	// 2584
	// 4181
}
