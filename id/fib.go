// Note: This isn't mine. And I do not remember where found it.
// I include it in this package as I consider it a fine beauty.
//
// The comment below is from the original author.

package id

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {

	a := make(chan int, 2)
	b := make(chan int, 2)
	c := make(chan int, 2)

	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	return out
}

// I like it because I only ever declare a single integer variable, and the rest goes through chans.
// This solution is inspired by Haskell's concise "fib = 0:zipWith (+) fib (1:fib)".

// ===========================================================================

// Fib returns the first N Fibonacci numbers
// starting with zero.
func Fib(N int) <-chan int {
	cha := make(chan int)
	go func(cha chan<- int) {
		x := fib()
		for i := 0; i < N; i++ {
			cha <- <-x
		}
		close(cha)
	}(cha)
	return cha
}

// ===========================================================================
