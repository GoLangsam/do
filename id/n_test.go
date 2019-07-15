// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/id/n_test.go"

package id_test

import (
	"fmt"
	"testing"

	"github.com/GoLangsam/do/id"
)

// ===========================================================================
// stolen from "github.com/bradfitz/iter"

func ExampleN() {
	for i := range id.N(4) {
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func TestAllocs(t *testing.T) {
	var x []struct{}
	allocs := testing.AllocsPerRun(500, func() {
		x = id.N(1e9)
		_ = x
	})
	if allocs > 0.1 {
		t.Errorf("allocs = %v", allocs)
	}
}

// ===========================================================================
