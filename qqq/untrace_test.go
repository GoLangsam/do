package qqq

import (
	"log"
	"os"
)

func Example_UnTrace() {
	verbose = true // we use qqq
	log.SetOutput(os.Stdout)
	log.SetPrefix("")

	// Exploit the fact that arguments to deferred functions are evaluated when the defer executes.
	// The trace function can set up the argument to the untracing routine.

	b := func() {
		defer un(trace("b"))
		qqq("in b")
	}

	a := func() {
		defer un(trace("a"))
		qqq("in a")
		b()
	}

	a()
	// Output:
	//
	// 	 Trace => a 	 => enter
	// in a
	// 	 Trace => b 	 => enter
	// in b
	// 	 Trace => b 	 <= leave
	// 	 Trace => a 	 <= leave

}
