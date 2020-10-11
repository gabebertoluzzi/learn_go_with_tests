// https://golang.org/doc/code.html
// First statement in go src code must be a package name
// executable packages MUST always use package main
package main

// https://golang.org/cmd/gofmt/
// Formats Go programs
// Default processes standard input, then prints to standard output.
import "fmt"

// use consts to improve performance with variables
const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const dutch = "Dutch"
const dutchHelloPrefix = "Hallo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case dutch:
		prefix = dutchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
