// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/qqq/qqq.log.go"

package qqq

import "log"

var se_ = log.Print
var see = log.Println
var sef = log.Printf

var di_ = log.Panic
var die = log.Panicln
var dif = log.Panicf

var verbose bool

func qq_(args ...interface{}) {
	if verbose {
		se_(args...)
	}
}

func qqq(args ...interface{}) {
	if verbose {
		see(args...)
	}
}

func qqf(format string, args ...interface{}) {
	if verbose {
		sef(format, args...)
	}
}
