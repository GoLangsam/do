// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package ats (= any to string) provides functions to Get a string from 'anything' (= interface{})
//
// ats observes different conventions of 'things' (=objects) to do so:
//  stringer: String() - fmt.Stringer & friends
//  string:   string   - builtin type
//  namer:    Name()   - filepath.File & .FileInfo, text/template.Template ...
//  geter:    Get()    - do/Value
//  ider:     Id()     - ...
//
// The different Get-functions just use a different sequence of attempts
// to obtain a meaningful string from interface{}, and return the best/first,
// or -as a last resort- a constant string such as "<not a Xyz>"
package ats // anything to a string

// I love to be friendly - thus: I observe different popular API's!
//  to convert anything to a meaningful text string:
//
// Note: this interface is exposed for godoc - only ;-)
type Friendly interface {
	String() string // fmt.Stringer & friends
	Name() string   // filepath.File & .FileInfo, text/template.Template ...
	Get() string    // do/Value
	Id() string     // ...
}

// an internal type for any observed interface

type stringer string

func (s stringer) String() string {
	return string(s)
}

type namer string

func (n namer) Name() string {
	return string(n)
}

type geter string

func (g geter) Get() string {
	return string(g)
}

// Note: another popular use is as getany interface{ Get() interface{} }.
// Intentionally we do not support this, as it may lead to infinite recursion.
// A getany may return a getany, which does not implement any of the other methods ...

type ider string

func (i ider) Id() string {
	return string(i)
}

// Get-Functions - content oriented

// GetString: string, String, Get, Name, Id, "<not a String>"
func GetString(any interface{}) string {

	switch typ := any.(type) {
	case string:
		return typ
	case stringer:
		return typ.String()
	case geter:
		return typ.Get()
	case namer:
		return typ.Name()
	case ider:
		return typ.Id()
	default:
		return "<not a String>"
	}
}

// GetGet: Get, String, string, Name, Id, "<not a Get>"
func GetGet(any interface{}) string {

	switch typ := any.(type) {
	case geter:
		return typ.Get()
	case stringer:
		return typ.String()
	case string:
		return typ
	case namer:
		return typ.Name()
	case ider:
		return typ.Id()
	default:
		return "<not a Get>"
	}
}

// Get-Functions - name/id oriented

// GetName: Name, Id, Get, String, string, "<not a Name>"
func GetName(any interface{}) string {

	switch typ := any.(type) {
	case namer:
		return typ.Name()
	case ider:
		return typ.Id()
	case geter:
		return typ.Get()
	case stringer:
		return typ.String()
	case string:
		return typ
	default:
		return "<not a Name>"
	}
}

// GetId: Id, Name, Get, String, string, "<not an Id>"
func GetId(any interface{}) string {

	switch typ := any.(type) {
	case ider:
		return typ.Id()
	case namer:
		return typ.Name()
	case geter:
		return typ.Get()
	case stringer:
		return typ.String()
	case string:
		return typ
	default:
		return "<not an Id>"
	}
}
