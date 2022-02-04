package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main(){
	
	// METHOD 1: unmarshalling json from different file
		var persons = getPersons()

		for i, _ := range persons {
			persons[i].Name = "\n this is name: " + persons[i].Name
		}

		fmt.Println("\nnew persons: ", persons)


	// METHOD 2: unmarshalling json in string format

		stringFrom := `{"full_name": "bradPitt", "age": 57, "gender":"male", "transport": {"car_name":"mercedes", "car_country":"germany","car_price":200000}}`
		// fmt.Printf("%T", stringFrom) 

		var bradPitt User 
		err := json.Unmarshal([]byte(stringFrom), &bradPitt)
		if err != nil {
			fmt.Println("no good")
		}
		fmt.Println("\nbradPitt: ", bradPitt)

		//MARSHALLING
		stringTo, err := json.Marshal(bradPitt)
		if err != nil {
			fmt.Print("error in marshalling data")
		}

		// fmt.Printf("%T",stringTo) //stringTo is type []uint8
		fmt.Printf("\n%s",stringTo) // print as string
		
	
	// METHOD 3: unmarshalling json using map
		
		var data map[string]interface{} //interface{} is a type that can be any value
		fmt.Println("\ntype of stringFrom: ", reflect.TypeOf(stringFrom))
		fmt.Println([]byte(stringFrom)) //converting the sTo to byte array and printing it
		
		err2 := json.Unmarshal([]byte(stringTo), &data)
		
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("\ndata: ", data) //converting the sTo to byte array and printing it
		fmt.Printf("\nTransport: ", data["Transport"])
		
		transport, ok := data["Transport"].(map[string]interface{})

		fmt.Println(transport)
		if ok {
			fmt.Println(transport["car_price"])
			} else {
			fmt.Println("\nfailed")
		}
		
}