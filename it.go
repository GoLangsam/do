// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do

// ===========================================================================

// It represents some action: do.It.
//
// The null value is useful: its Do() never does anything: it's a nop.
type It func()

// Do applies It iff It is not nil.
func (it *It) Do() {
	if *it != nil {
		(*it)()
	}
}

// ===========================================================================

// ItJoin returns a closure around given fs.
//
// Iff there are no fs, nil is returned, and
// iff there is only one fs, this single fs is returned.
//
// Evaluate the returned function
// by invoking its Do() or
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
				(&f).Do()
			}
		}
	}
}

// ===========================================================================

// Set sets all its as the new It action
// when the returned Option is applied.
func (it *It) Set(its ...It) Option {
	return func(any interface{}) Opt {
		prev := *it
		*it = ItJoin(its...)
		return func() Opt {
			return (*it).Set(prev)(any)
		}
	}
}

// Add adds all its before the existing It action
// when the returned Option is applied.
func (it *It) Add(its ...It) Option {
	if it == nil || *it == nil {
		return (*it).Set(its...)
	}
	return func(any interface{}) Opt {
		prev := *it
		*it = ItJoin(append(its, prev)...)
		return func() Opt {
			return (*it).Set(prev)(any)
		}
	}
}

// ===========================================================================
