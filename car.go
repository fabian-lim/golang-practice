package main

type Car struct {
	Brand string `json:"car_brand"`
	Type string	`json:"car_type"`
	From string `json:"car_country"`
	Price int `json:"car_price"`
}