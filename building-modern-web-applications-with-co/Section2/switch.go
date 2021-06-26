package main

import "log"

func main() {
	myVar := "cat2"

	switch myVar {
	case "cat":
		log.Println("myvar is set to cat")
	case "dog":
		log.Println("myvar is set to dog")

	default:
		log.Println("myvar is something else")

	}
}
