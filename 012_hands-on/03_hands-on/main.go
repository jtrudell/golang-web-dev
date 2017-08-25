package main

import (
	// "log"
	"os"
	"text/template"
)

type Place struct {
	Region string
}

type Hotel struct {
	Place
	Name, Address, City, Zip string
}

type HotelList struct {
	Name   string
	Hotels []Hotel
}

type Region struct {
	Name string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	hotel := Hotel{
		Place{"Southern"},
		"Holiday Inn",
		"555 East St",
		"Los Angeles",
		"90210",
	}

	hotel_two := Hotel{
		Place{"Northern"},
		"Best Western",
		"555 West St",
		"San Francisco",
		"90216",
	}

	some_hotels := []Hotel{hotel, hotel_two}
	data := HotelList{"California Hotels", some_hotels}

	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	// if err != nil {
	// 	log.Fatlln(err)
	// }
}
