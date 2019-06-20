// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.con/GoLangsam/container/oneway/list/form"

package do

// Option is a function which
// modifes something when applied to it
// and returns its own undo function (as Opt),
//  which, when applied, returns it's redo function (as Opt),
//  which, when applied, returns it's redo undo function (as Opt),
//  which, when applied, returns it's redo undo undo function (as Opt),
//  which, when applied, returns ...
//  (and so on: mind You: Opt is a self referential function).
//
// Hint: To apply multiple options at once, use do.Options(a, opts...).
//
// Note: Option will panic iff applied to the wrong type of object.
// (That is due to the known need of some type assertion in Go.)
//
// Hint: Provide Your own Options method and catch any panic in there.
//
// Note: This implementation was inspired by:
//   - http://commandcenter.blogspot.com.au/2014/01/self-referential-functions-and-design.html
//   - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
//   - https://www.calhoun.io/using-functional-options-instead-of-method-chaining-in-go/
// (Just: Undo is only supported for the last Option passed in these samples.)
//
// This implementation of do.Options(myType, myOpts...) provides full undo.
type Option func(interface{}) Opt

// Options applies every Option
// and returns its undo Opt-function
// which, when evaluated, fully
// restores the previous state.
func Options(a interface{}, options ...Option) Opt {

	prev := make([]Opt, 0, len(options))

	for i := range options { // apply each and remember it's undo
		prev = append(prev, options[i](a))
	}

	return func() Opt {
		return undo(prev...)
	}
}

// nopOpt is an Opt
// which returns a nopOpt
// (which returns a nopOpt)
// ((which returns a nopOpt))
// ((( ... )))
var nopOpt = func() Opt {
	return undo()
}

// undo applies the given Opt functions in reverse order
// and returns it's own undo Opt function.
//
// Thus, undo() without any Opt implements the nopOpt.
func undo(doit ...Opt) Opt {

	switch len(doit) {
	case 1:
		return doit[0]()
	case 0:
		return func() Opt { return undo() }
	default:
		prev := make([]Opt, 0, len(doit))
		for i := len(doit) - 1; i >= 0; i-- {
			prev = append(prev, doit[i]())
		}
		return func() Opt {
			return undo(prev...)
		}
	}
}
