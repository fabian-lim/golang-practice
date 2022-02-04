package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type person struct {
	Name string
	Age int
	Gender string
	Hobbies []string
	Address string
}



func getPersons()(persons []person) {
	fileBytes, err := ioutil.ReadFile("./people.json")

	if err != nil {
		panic(err)
	}

	fileContent := string(fileBytes) // have to convert the json file to string first
	fmt.Println("fileContent: ", fileContent)

	err  = json.Unmarshal(fileBytes, &persons)

	if err != nil {
		panic(err)
	}
	fmt.Println("print persons: ", persons) 
	return persons // returns json in array
}

	// personss := getPersons()

	// for i, _ := range persons {
	// 	persons[i].Name = persons[i].Name + "\n this is name"
	// }

	// fmt.Println(persons)