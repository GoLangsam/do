// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/id/c_test.go"

package id_test

import (
	"fmt"

	"github.com/GoLangsam/do/id"
)

// ===========================================================================
// finite channel

func ExampleC() {
	for i := range id.C(4) {
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

// ===========================================================================
