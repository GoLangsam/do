// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ats (= any to string) provides functions to Get a string from 'anything' (= interface{})
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

import (
	"github.com/golangsam/do/ami"
)

// Friendly - my interface - exposed for godoc - only ;-)
//
// I love to be friendly - thus: I observe different popular API's!
//  to convert anything to a meaningful text string:
type Friendly interface {
	String() string // fmt.Stringer & friends
	Name() string   // filepath.File & .FileInfo, text/template.Template ...
	Get() string    // do/Value
	Id() string     // ...
}

// an internal type for any observed interface
type stringer interface {
	String() string
}
type namer interface {
	Name() string
}
type geter interface {
	Get() string
}
type ider interface {
	Id() string
}

// Note: another popular use is as getany interface{ Get() interface{} }.
// Intentionally we do not support this, as it may lead to infinite recursion.
// A getany may return a getany, which does not implement any of the other methods ...

func notOk(any interface{}, typ string) string {
	if any == nil {
		return "" //	return "<ATS: Cannot get '" + typ + "' from '<nil>'!>"
	}
	return "<ATS: Cannot get '" + typ + "' from '" + ami.TypeName(any) + "'>"
}

// Get-Functions - content oriented

// GetString - order: string, String, Get, Name, Id, "<not a String>"
func GetString(any interface{}) string {

	if s, ok := any.(string); ok {
		return s
	}
	if s, ok := any.(stringer); ok {
		return s.String()
	}
	if s, ok := any.(geter); ok {
		return s.Get()
	}
	if s, ok := any.(namer); ok {
		return s.Name()
	}
	if s, ok := any.(ider); ok {
		return s.Id()
	}
	return notOk(any, "string")
}

// GetGet - order: Get, String, string, Name, Id, "<not a Get>"
func GetGet(any interface{}) string {

	if s, ok := any.(geter); ok {
		return s.Get()
	}
	if s, ok := any.(stringer); ok {
		return s.String()
	}
	if s, ok := any.(string); ok {
		return s
	}
	if s, ok := any.(namer); ok {
		return s.Name()
	}
	if s, ok := any.(ider); ok {
		return s.Id()
	}
	return notOk(any, "Get")
}

// Get-Functions - name/id oriented

// GetName - order: Name, Id, Get, String, string, "<not a Name>"
func GetName(any interface{}) string {

	if s, ok := any.(namer); ok {
		return s.Name()
	}
	if s, ok := any.(ider); ok {
		return s.Id()
	}
	if s, ok := any.(geter); ok {
		return s.Get()
	}
	if s, ok := any.(stringer); ok {
		return s.String()
	}
	if s, ok := any.(string); ok {
		return s
	}
	return notOk(any, "Name")
}

// GetId - order: Id, Name, Get, String, string, "<not an Id>"
func GetId(any interface{}) string {

	if s, ok := any.(ider); ok {
		return s.Id()
	}
	if s, ok := any.(namer); ok {
		return s.Name()
	}
	if s, ok := any.(geter); ok {
		return s.Get()
	}
	if s, ok := any.(stringer); ok {
		return s.String()
	}
	if s, ok := any.(string); ok {
		return s
	}
	return notOk(any, "Id")
}
