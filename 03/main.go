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
	next1, _ := iter.Pull(count(5))
	for {
		i, ok := next1()
		if !ok {
			break
		}

		fmt.Println("value: ", i)
	}

	next2, stop2 := iter.Pull2(squares(5))
	for {
		a, b, ok := next2()
		if !ok {
			break
		}

		fmt.Println("value: ", a, b)

		// 3 以降、 next2() から値が取得できなくなる
		if a >= 3 {
			fmt.Println("stop!!")
			stop2()
		}
	}
}

func main() {
	doSomething()
}
