package main

import "fmt"

func main() {

	defer myFunc()()
	fmt.Printf("hello!")
}

func myFunc() func() {
	fmt.Printf("juce")
	return func() {
		fmt.Println("myFunc")
	}
}
