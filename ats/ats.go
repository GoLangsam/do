// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package ats (= any to string) provides functions to Get a string from 'anything' (= interface{})
//
// ats observes different conventions of 'things' (=objects) to do so:
//  stringer: String()
//  namer:    Name()
//  geter:    Get()
//  ider:     Id()
//  string:   string
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
	Name() string   // filepath.File, filepath.FileInfo, text/template.Template ...
	Get() string    //
	Id() string     //
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

type ider string

func (i ider) Id() string {
	return string(i)
}

// Get-Functions - content oriented

// GetString: string, String, Get, Name, Id, "<not a String>"
func GetString(i interface{}) string {

	switch t := i.(type) {
	case string:
		return t
	case stringer:
		return t.String()
	case geter:
		return t.Get()
	case namer:
		return t.Name()
	case ider:
		return t.Id()
	default:
		return "<not a String>"
	}
}

// GetGet: Get, String, string, Name, Id, "<not a Get>"
func GetGet(i interface{}) string {

	switch t := i.(type) {
	case geter:
		return t.Get()
	case stringer:
		return t.String()
	case string:
		return t
	case namer:
		return t.Name()
	case ider:
		return t.Id()
	default:
		return "<not a Get>"
	}
}

// Get-Functions - name/id oriented

// GetName: Name, Id, Get, String, string, "<not a Name>"
func GetName(i interface{}) string {

	switch t := i.(type) {
	case namer:
		return t.Name()
	case ider:
		return t.Id()
	case geter:
		return t.Get()
	case stringer:
		return t.String()
	case string:
		return t
	default:
		return "<not a Name>"
	}
}

// GetId: Id, Name, Get, String, string, "<not an Id>"
func GetId(i interface{}) string {

	switch t := i.(type) {
	case ider:
		return t.Id()
	case namer:
		return t.Name()
	case geter:
		return t.Get()
	case stringer:
		return t.String()
	case string:
		return t
	default:
		return "<not an Id>"
	}
}
