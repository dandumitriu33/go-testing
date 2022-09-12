// First package in the testing series.
package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const romanian = "Romanian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const romanianHelloPrefix = "Salutare, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case romanian:
		prefix = romanianHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Christian", ""))
}
