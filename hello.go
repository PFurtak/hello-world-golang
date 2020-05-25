package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	foo()

	fmt.Println("Another line")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("foo func")
}

func bar() {
	fmt.Println("Program complete")
}
