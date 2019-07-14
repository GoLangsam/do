// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id

// Reluctant doubling

// S(1), S(2), ... = 1 1 2 1 1 2 4 1 1 2 1 1 2 4 8 1 1 2 1 1 2 4 1 1 2 ...

// S(n+1) = 2 * S(n) iff S(n) has already occured an even number of times, else
// S(n+1) = 1

// M. Luby, A. Sinclair, and D. Zuckerman [Information Proc. Letters 47 (1993), page 173-180,

func RelucantDoubles() <-chan int {
	cha := make(chan int)
	go func() {
		for u, v := 1, 1; ; {
			cha <- v
			if (u & -u) == v {
				u, v = u+1, 1
			} else {
				u, v = u, 2*v
			}
		}
	}()
	return cha
}

// ===========================================================================

// RelucantDouble returns the first N relucant doubled numbers
// starting with zero.
func RelucantDouble(N int) <-chan int {
	cha := make(chan int)
	go func(cha chan<- int) {
		x := RelucantDoubles()
		for i := 0; i < N; i++ {
			cha <- <-x
		}
		close(cha)
	}(cha)
	return cha
}

// ===========================================================================
