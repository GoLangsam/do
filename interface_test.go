// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do_test

import (
	"github.com/GoLangsam/do"
)

// ===========================================================================

func Example_interfaces() {

	// doer represents anyone who can apply some action
	// - usually as a closure around itself.
	type doer interface {
		Do()
	}

	// Iter represents anyone who can apply some action
	// - usually as a closure around itself.
	type Iter interface {
		doer
		Set(its ...do.It) do.Option
		Add(its ...do.It) do.Option
	}

	// Iffer represents anyone who can apply some action iff.
	// - usually as a closure around itself.
	type Iffer interface {
		doer
		Set(its ...do.It) do.Option
		Add(its ...do.It) do.Option
	}

	// Booler represents anyone who can provide some boolean
	// - usually as a closure around itself.
	type booler interface {
		Do() bool
	}

	// Oker represents anyone who can provide some boolean
	// (by default: true)
	// - usually as a closure around itself.
	type Oker interface {
		booler
		Set(oks ...do.Ok) do.Option
		Add(oks ...do.Ok) do.Option
	}

	// Noker represents anyone who can provide some boolean
	// (by default: false)
	// - usually as a closure around itself.
	type Noker interface {
		booler
		Set(noks ...do.Nok) do.Option
		Add(noks ...do.Nok) do.Option
	}

	// Errer represents anyone who can provide some error
	// - usually as a closure around itself.
	type Errer interface {
		Do() error
		Set(errs ...do.Err) do.Option
		Add(errs ...do.Err) do.Option
	}

	// Opter represents anyone who can provide some option
	// - usually as a closure around itself.
	type Opter interface {
		Do() do.Opt
		Set(opts ...do.Opt) do.Option
		Add(opts ...do.Opt) do.Option
	}

	var doit do.It = func() { return }
	var ok do.Ok = func() bool { return true }
	var iff do.If
	var nok do.Nok = func() bool { return false }
	var err do.Err = func() error { return nil }
	var opt do.Opt

	var _ doer = &doit
	var _ Iter = &doit
	var _ Iffer = &iff
	var _ Oker = &ok
	var _ Noker = &nok
	var _ Errer = &err
	var _ Opter = &opt

}

// ===========================================================================
