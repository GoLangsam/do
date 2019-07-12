// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id

// ===========================================================================

// I allows to range over the first N ordinal numbers:
//    for i := range I(10) {
//        fmt.Println(i)
//    }
// ... will print 1 to 10, inclusive.
func I(N int) <-chan int {
	Is := make(chan int)
	go func(Is chan<- int, N int) {
		defer close(Is)
		for i := 1; i < N+1; i++ {
			Is <- i
		}
	}(Is, N)
	return Is
}

// ===========================================================================
