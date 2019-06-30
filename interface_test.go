// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do_test

import (
	"github.com/GoLangsam/do"
)

// ===========================================================================

func Example_interfaces() {

	// Doer represents anyone who can apply some action
	// - usually as a closure around itself.
	type Doer interface {
		Do()
	}

	// Iffer represents anyone who can apply some action iff.
	// - usually as a closure around itself.
	type Iffer interface {
		Do()
	}

	// Oker represents anyone who can provide some boolean
	// (by default: true)
	// - usually as a closure around itself.
	type Oker interface {
		Do() bool
	}

	// Noker represents anyone who can provide some boolean
	// (by default: false)
	// - usually as a closure around itself.
	type Noker interface {
		Do() bool
	}

	// Errer represents anyone who can provide some error
	// - usually as a closure around itself.
	type Errer interface {
		Do() error
	}

	// Opter represents anyone who can provide some option
	// - usually as a closure around itself.
	type Opter interface {
		Do() do.Opt
	}

	var doit do.It = func() { return }
	var ok do.Ok = func() bool { return true }
	var iff do.If
	var nok do.Nok = func() bool { return false }
	var err do.Err = func() error { return nil }
	var opt do.Opt

	var _ Doer = &doit
	var _ Iffer = &iff
	var _ Oker = &ok
	var _ Noker = &nok
	var _ Errer = &err
	var _ Opter = &opt

}

// ===========================================================================
