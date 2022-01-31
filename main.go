package main

import (
	"fmt"
)

func main(){
	
	persons := getPersons()

	for i, _ := range persons {
		persons[i].Name = persons[i].Name + "\n this is name"
	}

	fmt.Println(persons)
}