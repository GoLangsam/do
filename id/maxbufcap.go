// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id

// ===========================================================================

// MaxBufCap defines the maximal capacity for buffered channels.
//
// Some returned channels are buffered
// in order to reduce the time used
// for context-switching among go routines.
var MaxBufCap = 1024

func maxBufCap(N int) int {
	if MaxBufCap < 1 {
		return 0
	}
	if N > MaxBufCap {
		return MaxBufCap
	}
	return N
}

// ===========================================================================
