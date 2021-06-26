package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Peter",
		LastName:  "Wirdemo",
	}

	myMap["me"] = me

	log.Println("Me is", myMap["me"].FirstName)
}
