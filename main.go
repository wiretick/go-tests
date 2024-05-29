package main

import "fmt"

const englishIntro = "Hello"

func hello(name string) string {
	if name == "" {
		return englishIntro + " World"
	}

	return fmt.Sprintf("%s %s", englishIntro, name)
}

func main() {
	fmt.Println(hello("Someone"))
}
