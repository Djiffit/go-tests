package main

import "fmt"

// Hello to the person
func Hello(person string) string {
	return "Hello, " + person + "!"
}

func main() {
	fmt.Println(Hello("World"))
}
