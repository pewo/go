package main

import "log"

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Peter")
	mySlice = append(mySlice, "Lovisa")
	mySlice = append(mySlice, "Linn")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5}

	log.Println(numbers)
}
