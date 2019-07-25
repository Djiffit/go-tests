package main

import "fmt"

const helloPrefixEn = "Hello, "
const helloPrefixSp = "Hola, "
const helloPrefixFr = "Bonjour, "
const spanish = "Spanish"
const french = "French"

// Hello to the person
func Hello(person string, language string) string {
	if person == "" {
		person = "World"
	}

	prefix := greetingPrefix(language)
	return prefix + person
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSp
	case french:
		prefix = helloPrefixFr
	default:
		prefix = helloPrefixEn
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
