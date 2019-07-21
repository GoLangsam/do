// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id

import (
	"fmt"
)

// ===========================================================================

// IDs allows to range over the first N IDs:
//    for c := range IDs("Prefix-", 100) {
//        fmt.Println(i)
//    }
// ... will print Prefix-001 to Prefix-100, inclusive.
func IDs(prefix string, N int) <-chan string {
	cha := make(chan string, maxBufCap(N))
	go func(cha chan<- string, N int) {
		defer close(cha)
		var f = getFormatWidth(prefix, N)
		for i := 0; i < N; i++ {
			cha <- fmt.Sprintf(f, prefix, i+1)
		}
	}(cha, N)
	return cha
}

// ===========================================================================
