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
// and should returns it's own undo Opt.
//
// The null value is useful: its Do() is a nop and returns a nop.Opt.
type Opt func() Opt

// Do applies Opt iff Opt is not nil.
func (a Opt) Do() Opt {
	if a != nil {
		return a()
	}
	return nopOpt
}

// ===========================================================================

// Join returns a closure around given fs.
//
// Iff there are no fs, a nop.Opt is returned, and
// iff there is only one fs, this single fs is returned.
//
// Evaluate the returned function
// by invoking it's Do() or
// by invoking it directly, iff not nil.
func (do *Opt) Join(fs ...Opt) Opt {
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

// WrapIt returns an Opt function
// which Do()es the Join of the given fs
// and returns its default, namely it's undo Opt,
// upon evaluation.
//
// Evaluate the returned function
// by invoking it's Do() or
// by invoking it directly, iff not nil.
//
// WrapIt may look like a convenient method.
//
// Just beware:
//  WrapIt violates the option contract:
//  no working undo Opt is returned
//  only a nop.Opt.
func (do *Opt) WrapIt(fs ...It) Opt {
	return func() Opt {
		var it It
		it.Join(fs...).Do()
		return nopOpt
	}
}

// ===========================================================================
