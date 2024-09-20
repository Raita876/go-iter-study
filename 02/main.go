package main

import (
	"fmt"
	"iter"
)

func count(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			if !yield(i) {
				break
			}
		}
	}
}

func squares(n int) iter.Seq2[int, int64] {
	return func(yield func(int, int64) bool) {
		for i := range n {
			if !yield(i, int64(i)*int64(i)) {
				break
			}
		}
	}
}

func doSomething() {
	for i := range count(5) {
		fmt.Println("value: ", i)
	}

	for a, b := range squares(5) {
		fmt.Println("value: ", a, b)
	}
}

func main() {
	doSomething()
}
