package main

import "fmt"

const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := englishHelloPrefix

	french := ""
	switch language {
	case french:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("fciasth", "french"))
}
