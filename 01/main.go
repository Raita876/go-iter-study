package main

import "fmt"

func fA(yield func() bool) {
	for range 3 {
		if !yield() {
			return
		}
	}
}

func fB(yield func(int) bool) {
	for i := range 3 {
		if !yield(i) {
			return
		}

	}
}

func fC(yield func(int, int) bool) {
	for i := range 5 {
		if !yield(i, i+5) {
			return
		}

	}
}

func doSomething() {
	for range fA {
		fmt.Println("Hello World!")
	}

	for i := range fB {
		fmt.Println("value: ", i)
	}

	for a, b := range fC {
		fmt.Println("value: ", a, b)
	}
}

func main() {
	doSomething()
}
