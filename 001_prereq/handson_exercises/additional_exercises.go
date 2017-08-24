package main

import (
	"fmt"
)

func main() {
	// initialize a SLICE of int using a composite literal;
	// print out the slice; range over the slice printing out just the index;
	// range over the slice printing out both the index and the value
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	for i, _ := range x {
		fmt.Println(i)
	}

	for i, val := range x {
		fmt.Println(i, val)
	}

	// Initialize a MAP using a composite literal where the key is a string and the value is an int;
	// print out the map; range over the map printing out just the key;
	// range over the map printing out both the key and the value

	y := map[string]int{"hey": 1, "later": 2, "yup": 3}
	fmt.Println(y)

	for key, _ := range y {
		fmt.Println(key)
	}

	for key, value := range y {
		fmt.Println(key, value)
	}

	// Create a new type: vehicle. The underlying type is a struct. The fields: doors, color.
	// Create two new types: truck & sedan. The underlying type of each of these new types is a struct.
	// Embed the “vehicle” type in both truck & sedan. Give truck the field “fourWheel” which will be set to bool.
	// Give sedan the field “luxury” which will be set to bool.

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	// Using the vehicle, truck, and sedan structs: using a composite literal, create a value of type truck and assign values to the fields;
	// using a composite literal, create a value of type sedan and assign values to the fields.
	// Print out each of these values.Print out a single field from each of these values.

	f := truck{vehicle{2, "red"}, true}
	g := sedan{vehicle{4, "black"}, false}
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println("The truck has", f.doors, "doors. Is it four wheel drive?", f.fourWheel)
	fmt.Println("The sedan is", g.color+".", "Is it luxury?", g.luxury)

}
