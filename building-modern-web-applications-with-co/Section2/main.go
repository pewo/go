package main

import "log"

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "bellut"
	log.Println(myMap["dog"])
}
