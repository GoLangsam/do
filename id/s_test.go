// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/id/s_test.go"

package id_test

import (
	"fmt"

	"github.com/GoLangsam/do/id"
)

// ===========================================================================
// slice

func ExampleS() {
	for _, i := range id.S("ID-", 4) {
		fmt.Println(i)
	}
	// Output:
	// ID-1
	// ID-2
	// ID-3
	// ID-4
}

// ===========================================================================
