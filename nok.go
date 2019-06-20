// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do

// ===========================================================================

// Nok represents some condition
// which is false by default.
//
// The null value is useful: its Do() never returns true.
type Nok func() bool

// Do applies Ok iff Ok is not nil,
// and makes Ok an Oker.
func (a Nok) Do() bool {
	if a != nil {
		return a()
	}
	return false
}

// ===========================================================================

// Join returns a closure around given fs.
//
// Iff there are no fs, nil is returned, and
// iff there is only one fs, this single fs is returned.
//
// Evaluate the returned function
// by invoking it's Do() or
// by invoking it directly, iff not nil.
func (do *Nok) Join(fs ...Nok) Nok {
	switch len(fs) {
	case 0:
		return nil
	case 1:
		return fs[0]
	default:
		return func() bool {
			for _, f := range fs {
				if f.Do() {
					return true
				}
			}
			return false
		}
	}
}

// ===========================================================================

// WrapIt returns a Nok function
// which Do()es the Join of the given fs
// and returns its default, namely: `false`,
// upon evaluation.
//
// Evaluate the returned function
// by invoking it's Do() or
// by invoking it directly, iff not nil.
//
// WrapIt is a convenient method.
func (do *Nok) WrapIt(fs ...It) Nok {
	return func() bool {
		var it It
		it.Join(fs...).Do()
		return false
	}
}

// ===========================================================================
