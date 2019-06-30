// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do

// ===========================================================================

// It represents some action: do.It.
//
// The null value is useful: its Do() is a nop.
type It func()

// Do applies It iff It is not nil,
// and makes It a Doer.
func (a It) Do() {
	if a != nil {
		a()
	}
}

// ===========================================================================

// ItJoin returns a closure around given fs.
//
// Iff there are no fs, nil is returned, and
// iff there is only one fs, this single fs is returned.
//
// Evaluate the returned function
// by invoking it's Do() or
// by invoking it directly, iff not nil.
func ItJoin(fs ...It) It {
	switch len(fs) {
	case 0:
		return nil
	case 1:
		return fs[0]
	default:
		return func() {
			for _, f := range fs {
				f.Do()
			}
		}
	}
}

// ===========================================================================
