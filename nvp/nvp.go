// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package nvp provides helper functions for any KeyValuePair,
// which satisfies the K/V/GetV = Key/Value interface.
//
// TODO: support other popular naming conventions, e.g.
//  Key Value
//  Id Is
//
// Some functions help to utilise the recursive nature of an arbitrary KeyValuePair.
package nvp

import (
	"path"
)

// Friendly - my interface - exposed for godoc - only ;-)
//
// Note: You may like to use my own kind as stuff ;-)
//  Thus: build a path to where I hide - KeyValuePairly!
//  just: Using Named() in a loop may not get You much :-(
//  just the same name over and over ... and over again ...
//  thus: better use me.Into(key string) to build Your path ;-)
//
// How to use? Easy:
//  Use Into to hide the treasure under another name
//  Use Leaf to retrieve the treasure
//  Use NameS or Path - they'll tell You where the treasure is hidden
//  ... just in case You forgot ;-)
type Friendly interface {
	NamePath() string                // return the names leading to the hidden treasure as a (cleaned!) "path"
	Leaf(p KeyValuePair) interface{} // return the treasure hidden deep inside KeyValuePair's
	NameS() []string                 // return the names leading to the hidden treasure as a slice of strings
}

// var _ Friendly = New("Interface satisfied? :-)")

// KeyValuePair - this interface allows me to recurse
// Note: this interface is exposed for godoc - only ;-)
type KeyValuePair interface {
	K() string
	V() string
	GetV() interface{}
}

// Leaf returns a new TagAny named "key" and containing d
func Leaf(p KeyValuePair) interface{} {

	var c interface{} = KeyValuePair(p)
	for {
		if x, ok := c.(KeyValuePair); ok {
			c = x.GetV() // try again - with my value ...
		} else {
			break
		}
	}
	return c
}

// NameS returns the names leading to the hidden treasure as a slice of strings
func NameS(p KeyValuePair) []string {

	var n []string

	var c interface{} = KeyValuePair(p)
	for {
		if x, ok := c.(KeyValuePair); ok {
			n = append(n, x.K()) // remember my name, and ...
			c = x.GetV()         // try again - with my value ...
		} else {
			break
		}
	}
	return n
}

// NamePath returns the names leading to the hidden treasure as a "path"
//  Note: as the "path" is cleaned, it may not lead You back to the treasure!
func NamePath(p KeyValuePair) string {

	var n string

	var c interface{} = KeyValuePair(p)
	for {
		if x, ok := c.(KeyValuePair); ok {
			n = path.Join(n, x.K()) // remember my name, and ...
			c = x.GetV()            // try again - with my value ...
		} else {
			break
		}
	}
	return n
}

// TODO: try to use this functional approach. What to do for leaf - returns interface{}
func nameS(p KeyValuePair) []string {
	var n []string
	var f = func(p KeyValuePair, acc []string) {
		n = append(n, p.K()) // remember my name, and ...
	}
	eval(p, f, n)
	return n
}

func namepath(p KeyValuePair) string {
	var n []string = make([]string, 1)
	var f = func(p KeyValuePair, acc []string) {
		n[0] = path.Join(n[0], p.K()) // remember my name, and ...
	}
	eval(p, f, n)
	return n[0]
}

func eval(p KeyValuePair, f func(p KeyValuePair, acc []string), res []string) {
	var c interface{} = KeyValuePair(p)
	for {
		if x, ok := c.(KeyValuePair); ok {
			f(x, res)    // apply function to me, and ...
			c = x.GetV() // try again - with my value ...
		} else {
			break
		}
	}
}
