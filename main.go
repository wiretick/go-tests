package main

import "fmt"

const (
	langNo         = "Norwegian"
	langFr         = "French"
	englishIntro   = "Hello"
	frenchIntro    = "Bonjour"
	norwegianIntro = "Hei"
	defaultName    = "World"
)

func hello(name, lang string) string {
	var intro string
	switch lang {
	case langNo:
		intro = norwegianIntro
	case langFr:
		intro = frenchIntro
	default:
		intro = englishIntro
	}

	if name == "" {
		name = defaultName
	}

	return fmt.Sprintf("%s %s", intro, name)
}

func main() {
	fmt.Println(hello("Someone", ""))
}
