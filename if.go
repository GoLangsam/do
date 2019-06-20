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
// The null value is useful: its Do() is a nop.
type If struct {
	It
	If bool
}

// Do applies It iff If and It is not nil,
// and makes If a Doer.
func (a If) Do() {
	if a.If && a.It != nil {
		a.It()
	}
}

// ===========================================================================
