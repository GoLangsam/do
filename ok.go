// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do

// ===========================================================================

// Ok represents some condition
// which is true by default.
//
// The null value is useful: its Do() never returns false.
type Ok func() bool

// Do applies Ok iff Ok is not nil,
// and makes Ok an Oker.
func (ok *Ok) Do() bool {
	if *ok != nil {
		return (*ok)()
	}
	return true
}

// ===========================================================================

// OkJoin returns a closure around given fs.
//
// Iff there are no fs, nil is returned, and
// iff there is only one fs, this single fs is returned.
//
// Evaluate the returned function
// by invoking its Do() or
// by invoking it directly, iff not nil.
//
// Note: Order matters - evaluation terminates on first exceptional (non-default) result.
func OkJoin(fs ...Ok) Ok {
	switch len(fs) {
	case 0:
		return nil
	case 1:
		return fs[0]
	default:
		return func() bool {
			for _, f := range fs {
				if !(&f).Do() {
					return false
				}
			}
			return true
		}
	}
}

// ===========================================================================

// Set sets all oks as the new ok
// when the returned Option is applied.
func (ok *Ok) Set(oks ...Ok) Option {
	return func(any interface{}) Opt {
		prev := *ok
		*ok = OkJoin(oks...)
		return func() Opt {
			return (*ok).Set(prev)(any)
		}
	}
}

// Add adds all oks before the existing ok
// when the returned Option is applied.
func (ok *Ok) Add(oks ...Ok) Option {
	if ok == nil || *ok == nil {
		return (*ok).Set(oks...)
	}
	return func(any interface{}) Opt {
		prev := *ok
		*ok = OkJoin(append(oks, prev)...)
		return func() Opt {
			return (*ok).Set(prev)(any)
		}
	}
}

// ===========================================================================

// OkIt returns an Ok function
// which Do()es the Join of the given its
// and returns the default, namely: `true`,
// upon evaluation.
//
// Evaluate the returned function
// by invoking it's Do() or
// by invoking it directly, iff not nil.
//
// OkIt is a convenient wrapper.
func OkIt(its ...It) Ok {
	return func() bool {
		it := ItJoin(its...)
		(&it).Do()
		return true
	}
}

// ===========================================================================
