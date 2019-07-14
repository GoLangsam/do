// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package id provides a few simple iterables
// useful to range over and/or to produce some testdata such as:
//
//  some finite read-only channel:
//  - I(N) to range over the first N ordinal numbers 1...N
//  - C(N) to range over the first N+1 cardinal numbers 0...N
//  - Names(prefix, N) to range over N size-adjusted prefixed IDs.
//
//  some slice:
//  - S(prefix, N) => a slice of N size-adjusted prefixed IDs.
//  - N(N) a slice to range over the first N offset numbers 0...N-1
//    (A fine joke - 'stolen' from "github.com/bradfitz/iter")
//
// A prefixed ID is a string composed of a given prefix string
// followed by a dash and a 0-padded increasing number-string,
// e.g. "ID-01" ... "ID-13".
// For an empty prefix, right-adjusted number-strings are produced.
//
// Further, there are
//  - reluctantly doubled numbers (as found in TAOCP)
//  - Fibonacci numbers
//    (with an interesting channel-based implementation)
// Finite and non-terminating iterators are supplied.
package id
