package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	id      int
	name    string
	degrees []string
}

type Animal struct {
	name string `required_max:"100"`
}

type Bird struct {
	Animal
	speed int
}

func main() {
	// map	-	reference type
	aMap := make(map[string]int) // map of string (key) -> int (value)
	fmt.Printf("%v\n", aMap)

	bMap := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Printf("%v\n", bMap["one"])

	bMap["three"] = 3
	fmt.Printf("%v\n", bMap)

	value3, ok3 := bMap["three"]
	fmt.Printf("value: %v, exists: %v\n", value3, ok3)

	value4, ok4 := bMap["four"]
	fmt.Printf("value: %v, exists: %v\n", value4, ok4)

	// struct	-	value type
	aDoctor := Doctor{
		id:   1,
		name: "John Doe",
		degrees: []string{
			"MBBS",
			"FCPS",
		},
	}
	fmt.Printf("%v, %T\n", aDoctor, aDoctor)

	bDoctor := aDoctor // create copy
	fmt.Printf("%v, %T\n", bDoctor, bDoctor)

	cDoctor := &aDoctor // creating reference
	fmt.Printf("%v, %T\n", cDoctor, cDoctor)

	anotherDoctor := struct { // anonymous struct
		name string
	}{"Jane Doe"}
	fmt.Printf("%v, %T\n", anotherDoctor, anotherDoctor)

	// Bird does not inherit Animal -> Bird is not an animal, rather Bird contains Animal attributes
	// go add syntactical sugar to access embedded objects like the attributes belong to Bird
	// this is called go composition
	aBird := Bird{}
	aBird.name = "Crow"
	aBird.speed = 55
	fmt.Printf("%v, %T\n", aBird, aBird)

	// single line declaration and assignment has to consider embedded structs
	bBird := Bird{
		Animal: Animal{
			name: "Caccoo",
		},
		speed: 52,
	}
	fmt.Printf("%v, %T\n", bBird, bBird)

	// tags provide additional data about the field
	// tags have to be parsed and make meaning out of
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}
