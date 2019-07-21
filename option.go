// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
// (This is due to the known need of Go to assert the type dynamically.)
//
// Hint: Provide Your own Options method and catch any panic in there.
//
// Note: This implementation was inspired by:
//   - http://commandcenter.blogspot.com.au/2014/01/self-referential-functions-and-design.html
//   - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
//   - https://www.calhoun.io/using-functional-options-instead-of-method-chaining-in-go/
// (Just: in these samples Undo is only supported for the last Option passed.)
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

// ===========================================================================

// OptionJoin returns a closure around given fs.
//
// Iff there are no fs, a nop.Option is returned, and
// iff there is only one fs, this single fs is returned.
//
// Evaluate the returned function in order to apply its effect.
func OptionJoin(fs ...Option) Option {
	switch len(fs) {
	case 0:
		return func(any interface{}) Opt { return nopOpt }
	case 1:
		return fs[0]
	default:
		return func(any interface{}) Opt {
			prev := make([]Opt, 0, len(fs))
			for _, option := range fs {
				prev = append(prev, option(any))
			}
			return func() Opt {
				return undo(prev...)
			}
		}
	}
}

// ===========================================================================

// Set sets all options as the new option
// when the returned Option is applied.
func (option *Option) Set(options ...Option) Option {
	return func(any interface{}) Opt {
		prev := *option
		*option = OptionJoin(options...)
		return func() Opt {
			return (*option).Set(prev)(any)
		}
	}
}

// Add appends all options to the existing option
// when the returned Option is applied.
func (option *Option) Add(options ...Option) Option {
	if option == nil || *option == nil {
		return (*option).Set(options...)
	}
	return func(any interface{}) Opt {
		prev := *option
		*option = OptionJoin(append([]Option{prev}, options...)...)
		return func() Opt {
			return (*option).Set(prev)(any)
		}
	}
}

// ===========================================================================

// OptionIt returns an Option function
// which effects the Join of the given its
// when the returned Option is applied
// and returns the default, namely: its undo Opt.
//
// OptionIt may look like a convenient wrapper.
//
//  Just beware: OptIt violates the option contract!
//  No working undo Opt is returned - only a nop.Opt.
func OptionIt(its ...It) Option {
	return func(any interface{}) Opt {
		return OptIt(its...)()
	}
}

// ===========================================================================
