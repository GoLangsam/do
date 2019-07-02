// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do

// ===========================================================================

// Opt represents some option
// which can be taken.
//
// Opt is a self referential function
// obtained when some Option is applied
// and should return its own undo Opt.
//
// The null value is useful: its Do() never does anything
// as it is a nop and returns a nop.Opt.
type Opt func() Opt

// Do applies Opt iff Opt is not nil.
func (a Opt) Do() Opt {
	if a != nil {
		return a()
	}
	return nopOpt
}

// ===========================================================================

// OptJoin returns a closure around given fs.
//
// Iff there are no fs, a nop.Opt is returned, and
// iff there is only one fs, this single fs is returned.
//
// Evaluate the returned function
// by invoking its Do() or
// by invoking it directly, iff not nil.
func OptJoin(fs ...Opt) Opt {
	switch len(fs) {
	case 0:
		return nopOpt
	case 1:
		return fs[0]
	default:
		return undo(fs...)
	}
}

// ===========================================================================

// Set sets all opts as the new opt
// when the returned Option is applied.
func (opt *Opt) Set(opts ...Opt) Option {
	return func(any interface{}) Opt {
		prev := *opt
		*opt = OptJoin(opts...)
		return func() Opt {
			return (*opt).Set(prev)(any)
		}
	}
}

// Add appends all opts to the existing opt
// when the returned Option is applied.
func (opt *Opt) Add(opts ...Opt) Option {
	if opt == nil || *opt == nil {
		return (*opt).Set(opts...)
	}
	return func(any interface{}) Opt {
		prev := *opt
		*opt = OptJoin(append([]Opt{prev}, opts...)...)
		return func() Opt {
			return (*opt).Set(prev)(any)
		}
	}
}

// ===========================================================================

// OptIt returns an Opt function
// which Do()es the Join of the given its
// and returns the default, namely: its undo Opt,
// upon evaluation.
//
// Evaluate the returned function
// by invoking its Do() or
// by invoking it directly, iff not nil.
//
// OptIt may look like a convenient wrapper.
//
//  Just beware: OptIt violates the option contract!
//  No working undo Opt is returned - only a nop.Opt.
func OptIt(its ...It) Opt {
	return func() Opt {
		it := ItJoin(its...)
		(&it).Do()
		return nopOpt
	}
}

// ===========================================================================
