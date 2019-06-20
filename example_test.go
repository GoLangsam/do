// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package do_test

import (
	"fmt"

	"github.com/GoLangsam/container/oneway/list"
	"github.com/GoLangsam/do"
)

// Value returns a function which
// sets Value to v
// and returns it's undo Opt.
func Value(v interface{}) do.Option {
	return func(any interface{}) do.Opt {
		a := any.(*list.Element)
		prev := (*a).Value
		(*a).Value = v
		return func() do.Opt {
			return Value(prev)(a)
		}
	}
}

func ExampleOpt() {
	e := list.NewList("TestList", "Element One").Front()
	fmt.Println(e.Value) // Element One

	undo := Value(3)(e)
	fmt.Println(e.Value) // 3 (temporarily)

	redo := undo()
	fmt.Println(e.Value) // Element One (temporary setting undone)

	rere := redo()
	fmt.Println(e.Value) // 3 (undo undone)

	_ = rere()
	fmt.Println(e.Value) // Element One (temporary setting undone)

	// Output:
	// Element One
	// 3
	// Element One
	// 3
	// Element One
}

func ExampleOptions() {
	e := list.NewList("TestList", "Element One").Front()
	fmt.Println(e.Value) // Element One

	undo := do.Options(e, Value(3), Value("5"), Value(7))
	fmt.Println(e.Value) // 7 (temporarily)

	redo := undo()
	fmt.Println(e.Value) // Element One (temporary setting undone)

	_ = redo()
	fmt.Println(e.Value) // 7 (undo undone)

	// Output:
	// Element One
	// 7
	// Element One
	// 7
}

func Example_value() {

	// Value returns a function which
	// sets Value to v
	// and returns it's undo Opt.
	Value := func(v interface{}) do.Option {
		return func(any interface{}) do.Opt {
			a := any.(*list.Element)
			prev := a.Value
			a.Value = v
			return func() do.Opt {
				return Value(prev)(a)
			}
		}
	}

	e := list.NewList("TestList", "Element One").Front()

	setValue := func(e *list.Element, v interface{}) {
		// upon exit apply undo to restore original value while setting to new value v now via
		defer Value(v)(e)() // Note the triple evaluation.

		// ... do some stuff with Value being temporarily set to v.
		fmt.Println(e.Value) // Changed Value
	}

	fmt.Println(e.Value) // Original Value
	setValue(e, 5)
	fmt.Println(e.Value)

	// Output:
	// Element One
	// 5
	// Element One
}

func ExampleOptions_nochange() {
	e := list.NewList("TestList", "Element One").Front()

	undo := do.Options(e)
	fmt.Println(e.Value) // Element One (unchanged)

	redo := undo()
	fmt.Println(e.Value) // Element One (temporary no-change undone)

	_ = redo()
	fmt.Println(e.Value) // Element One (temporary no-change undo redone)

	// Output:
	// Element One
	// Element One
	// Element One
}
