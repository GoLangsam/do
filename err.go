// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do

// ===========================================================================

// Err represents some action
// which might go wrong (for some reason).
//
// The null value is useful: its Do() never returns a non-nil error.
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

// ErrJoin returns a closure around given fs.
//
// Iff there are no fs, nil is returned, and
// iff there is only one fs, this single fs is returned.
//
// Evaluate the returned function
// by invoking its Do() or
// by invoking it directly, iff not nil.
//
// Note: Order matters - evaluation terminates on first exceptional (non-default) result.
func ErrJoin(fs ...Err) Err {
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

// Set sets all errs as the new err
// when the returned Option is applied.
func (err *Err) Set(errs ...Err) Option {
	return func(any interface{}) Opt {
		prev := *err
		*err = ErrJoin(errs...)
		return func() Opt {
			return (*err).Set(prev)(any)
		}
	}
}

// Add appends all errs to the existing err
// when the returned Option is applied.
func (err *Err) Add(errs ...Err) Option {
	if err == nil || *err == nil {
		return (*err).Set(errs...)
	}
	return func(any interface{}) Opt {
		prev := *err
		*err = ErrJoin(append([]Err{prev}, errs...)...)
		return func() Opt {
			return (*err).Set(prev)(any)
		}
	}
}

// ===========================================================================

// ErrIt returns an Err function
// which Do()es the Join of the given its
// and returns the default, namely: `nil`,
// upon evaluation.
//
// Evaluate the returned function
// by invoking it's Do() or
// by invoking it directly, iff not nil.
//
// ErrIt is a convenient wrapper.
func ErrIt(its ...It) Err {
	return func() error {
		it := ItJoin(its...)
		(&it).Do()
		return nil
	}
}

// ===========================================================================
