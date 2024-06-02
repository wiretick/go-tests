package main

import (
	"fmt"
	"io"
	"os"
)

const (
	langNo = "Norwegian"
	langFr = "French"

	englishIntro   = "Hello"
	frenchIntro    = "Bonjour"
	norwegianIntro = "Hei"

	defaultName = "World"
)

func Hello(name, lang string) string {
	intro := generateIntro(lang)

	if name == "" {
		name = defaultName
	}

	return fmt.Sprintf("%s %s", intro, name)
}

func generateIntro(lang string) string {
	var intro string
	switch lang {
	case langNo:
		intro = norwegianIntro
	case langFr:
		intro = frenchIntro
	default:
		intro = englishIntro
	}

	return intro
}

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func main() {
	fmt.Println(Hello("Someone", ""))

	Greet(os.Stdout, "Sarah")
}
