// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package amy / `any meta info` - easy access to meta data
//
// I love to be informative - and even give metadata about my content anything
//  use TypeName to get the name of the type of my content
//  use TypePkgName to get the package name of the type of my content
//  use TypeString to get a 'nick-name' of the type of my content
//  use TypeKind to get the Kind of the type of my content ( int, struct, func, ...)
//  use TypeIsComparable ...
//  use TypeIsVariadic (only useful, if my TypeKind is a function)
// Note: I use the "reflect" package to obtain such metadata - as You may have guessed ;-)
package ami // any meta info

import (
	"reflect"
)

// TypeName returns the `reflect.TypeOf(any).Name()`
func TypeName(any interface{}) string {
	t := reflect.TypeOf(any)
	return t.Name()
}

// TypePkgPath returns the `reflect.TypeOf(any).PkgPath()`
func TypePkgPath(any interface{}) string {
	t := reflect.TypeOf(any)
	return t.PkgPath()
}

// TypeString returns the `reflect.TypeOf(any).String()`
func TypeString(any interface{}) string {
	t := reflect.TypeOf(any)
	return t.String()
}

// TypeKind returns the `reflect.TypeOf(any).Kind().String()`
func TypeKind(any interface{}) string {
	t := reflect.TypeOf(any)
	return t.Kind().String()
}

// TypeIsComparable returns the `reflect.TypeOf(any).Comparable()`
func TypeIsComparable(any interface{}) bool {
	t := reflect.TypeOf(any)
	return t.Comparable()
}

// TypeIsVariadic returns the `reflect.TypeOf(any).IsVariadic()` (and `false` for any non-Func)
func TypeIsVariadic(any interface{}) bool {
	t := reflect.TypeOf(any)
	if t.Kind() != reflect.Func {
		return false
	} else {
		return t.IsVariadic()
	}
}

// excerpts obtained via "godoc reflect" - temprarily copied for local reference

/*
 * These data structures are known to the compiler (../../cmd/internal/gc/reflect.go).
 * A few are known to ../runtime/type.go to convey to debuggers.
 * They are also known to ../runtime/type.go.
 */

/*
// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)
*/

// Methods applicable only to some types, depending on Kind.
// The methods allowed for each kind are:
//
//	Int*, Uint*, Float*, Complex*: Bits
//	Array: Elem, Len
//	Chan: ChanDir, Elem
//	Func: In, NumIn, Out, NumOut, IsVariadic.
//	Map: Key, Elem
//	Ptr: Elem
//	Slice: Elem
//	Struct: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField
