// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id_test

import (
	"fmt"

	"github.com/GoLangsam/do/id"
)

// ===========================================================================
// finite channel

func ExampleIDs() {
	for i := range id.IDs("ID-", 4) {
		fmt.Println(i)
	}
	// Output:
	// ID-1
	// ID-2
	// ID-3
	// ID-4
}

// ===========================================================================
