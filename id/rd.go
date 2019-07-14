// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id

// Reluctant doubling

// S(1), S(2), ... = 1 1 2 1 1 2 4 1 1 2 1 1 2 4 8 1 1 2 1 1 2 4 1 1 2 ...

// S(n+1) = 2 * S(n) iff S(n) has already occured an even number of times, else
// S(n+1) = 1

// M. Luby, A. Sinclair, and D. Zuckerman [Information Proc. Letters 47 (1993), page 173-180,

// nextReluctantDouble literally as found in "Donald Knuth, The Art of Computer Programming chapter 7.2.2.2"
func nextReluctantDouble(u, v int) (int, int) {
	if u & -u == v {
		return u + 1, 1
	} else {
		return u, 2 * v
	}

}

// ===========================================================================

// ReluctantDoubles sends reluctantly doubled numbers
// on the returned channel.
func ReluctantDoubles() <-chan int {
	cha := make(chan int)
	go func() {
		for u, v := 1, 1; ; u, v = nextReluctantDouble(u, v) {
			cha <- v
		}
	}()
	return cha
}

// ===========================================================================

// ReluctantDouble sends the first N relucantly doubled numbers
// on the returned channel.
func ReluctantDouble(N int) <-chan int {
	cha := make(chan int)
	var i int
	go func() {
		for u, v := 1, 1; i < N; u, v = nextReluctantDouble(u, v) {
			cha <- v
			i++
		}
		close(cha)
	}()
	return cha
}

// ===========================================================================
