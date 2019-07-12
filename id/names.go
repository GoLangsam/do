// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id

import (
	"fmt"
)

// ===========================================================================

// NameS allows to range over the first N Names:
//    for c := range NameS("Prefix-", 100) {
//        fmt.Println(i)
//    }
// ... will print Prefix-001 to Prefix-100, inclusive.
func NameS(prefix string, N int) <-chan string {
	names := make(chan string)
	go func(names chan<- string, N int) {
		defer close(names)
		var f = getFormatWidth(prefix, N)
		for i := 0; i < N; i++ {
			names <- fmt.Sprintf(f, prefix, i+1)
		}
	}(names, N)
	return names
}

// ===========================================================================
