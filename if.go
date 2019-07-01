// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do

// ===========================================================================

// If represents an opional action.
//
// It wraps It (an Action - pun intended)
// and a boolean If switch
// and it facilitates conditional invocation via its Do() method.
//
// Intended use is for conditional logging, counting etc.
//
// The null value is useful: its Do() never does anything, it's a nop.
type If struct {
	It
	If bool
}

// Do applies It iff If is true and It is not nil,
// and makes If a Doer.
func (a If) Do() {
	if a.If && a.It != nil {
		a.It()
	}
}

// ===========================================================================

// Iff makes iff the new If value
// when the returned Option is applied.
func (fn *If) Iff(iff bool) Option {
	return func(any interface{}) Opt {
		prev := (*fn).If
		(*fn).If = iff
		return func() Opt {
			return (*fn).Iff(prev)(any)
		}
	}
}

// ===========================================================================
