// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id

// ===========================================================================
// a little gem found @ "github.com/bradfitz/iter"

// N returns a slice of N+1 0-sized elements
// (all offset numbers up to and inclusiding N, so to say),
// suitable for ranging over:
//    for i := range N(10) {
//        fmt.Println(i)
//    }
// ... will print 0 to 10, inclusive.
//
// It does not cause any allocations.
func N(N int) []struct{} {
	return make([]struct{}, N+1)
}

// ===========================================================================
