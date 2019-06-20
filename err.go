// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do

// ===========================================================================

// Err represents some action
// which might go wrong (for some reason).
//
// The null value is useful: its Do() returns nil.
type Err func() error

// Do evaluates Err iff Err is not nil,
// and makes Err an Errer.
func (a Err) Do() error {
	if a != nil {
		return a()
	}
	return nil
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
func (do *Err) Join(fs ...Err) Err {
	switch len(fs) {
	case 0:
		return nil
	case 1:
		return fs[0]
	default:
		return func() error {
			var err error
			for _, f := range fs {
				err = f.Do()
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
}

// ===========================================================================

// WrapIt returns an Err function
// which Do()es the Join of the given fs
// and returns its default, namely: `nil`,
// upon evaluation.
//
// Evaluate the returned function
// by invoking it's Do() or
// by invoking it directly, iff not nil.
//
// WrapIt is a convenient method.
func (do *Err) WrapIt(fs ...It) Err {
	return func() error {
		var it It
		it.Join(fs...).Do()
		return nil
	}
}

// ===========================================================================
