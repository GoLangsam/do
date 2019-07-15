// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/id/i_test.go"

package id_test

import (
	"fmt"

	"github.com/GoLangsam/do/id"
)

// ===========================================================================
// finite channel

func ExampleI() {
	for i := range id.I(4) {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

// ===========================================================================
