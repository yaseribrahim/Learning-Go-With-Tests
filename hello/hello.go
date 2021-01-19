package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const arabic = "Arabic"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchhHelloPrefix = "Bonjour, "
const arabicHelloPrefix = "Salam, "

// Hello prints "Hello, World" or "Hello, {Person}"" in French, Spanish or English
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchhHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case arabic:
		prefix = arabicHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
