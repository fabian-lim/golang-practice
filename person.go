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

// 	person1 := person{
// 		Name: "fabian",
// 		Age: 27,
// 		Gender: "M",
// 		Hobbies: []string{"coding","tennis","netflix"},
// 		Address: "the vyne residence",
// 	}

// 	person2 := person{
// 		Name: "brad",
// 		Age: 50,
// 		Gender: "M",
// 		Hobbies: []string{"acting","drinking","pranks"},
// 		Address: "hollywood avenue",
// 	}

// 	return []person{ person1, person2 }

func getPersons()(persons []person) {
	fileBytes, err := ioutil.ReadFile("./people.json")

	if err != nil {
		panic(err)
	}

	fileContent := string(fileBytes)
	fmt.Println(fileContent)

	err  = json.Unmarshal(fileBytes, &persons)

	if err != nil {
		panic(err)
	}

	return persons
}