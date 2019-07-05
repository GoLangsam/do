// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/qqq/qqq.set.go"

package qqq

// SetVerbose controls the verbosity of this package.
//
// Iff true some detailed processing information may be logged.
func SetVerbose(v bool) {
	verbose = v
}

// GetVerbose reports the current verbosity of this package.
//
// Iff true some detailed processing information may be logged.
func GetVerbose() bool {
	return verbose
}
