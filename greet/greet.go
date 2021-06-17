package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const hindi = "Hindi"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const hindiHelloPrefix = "Namaste, "

// Hello returns a string.
// This function is public because it starts with a capital letter.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

// named return value 'prefix'. This creates a prefix variable in the function.
// Since its a string it has a default value of "".
// Since we used a named return value we did 'return' instead of 'return prefix'.
// This is a private function since it starts with a lowercase letter.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Raza", ""))
}
