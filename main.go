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
	if lang == langNo {
		intro = norwegianIntro
	} else if lang == langFr {
		intro = frenchIntro
	} else {
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
