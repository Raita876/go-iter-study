package main

import (
	"fmt"
	"iter"
	"time"
)

func countA(n int) []int {
	sl := make([]int, n)

	for i := range n {
		sl[i] = i
	}

	return sl
}

func countB(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			if !yield(i) {
				break
			}
		}
	}
}

func doSomething() {
	n := 1000000

	start1 := time.Now()
	for range countA(n) {
		// skip
	}
	since1 := time.Since(start1)
	fmt.Println(since1)

	start2 := time.Now()
	for range countB(n) {
		// skip
	}
	since2 := time.Since(start2)
	fmt.Println(since2)

}

func main() {
	doSomething()
}
