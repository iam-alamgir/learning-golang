package main

import "fmt"

func main() {
	// if-else
	// always use parenthesis
	if true {
		fmt.Println("It's your first IF")
	}

	states := map[string]int{
		"California": 123,
		"Florida":    456,
	}

	if pop, ok := states["Florida"]; ok {
		// pop is scoped to if block
		fmt.Println(pop)
	}

	// comparisons >, <, ==, !=, >=, <=
	guess := 50
	if guess < 50 {
		fmt.Println("less than 50")
	} else if guess > 50 {
		fmt.Println("greater than 50")
	} else {
		fmt.Println("equal to 50")
	}

	// switch
	// block is considered from case to case
	// breaks out after case block
	// fallthrough is manual and ignores case check
	// no case overlap when tagged syntax used

	// tagged
	switch guess {
	case 10:
		// single match
		fmt.Println("it's 10")
	case 20, 30, 40:
		// multiple match
		fmt.Println("it's 20, 30 or 40")
	default:
		// when everything fails
		fmt.Println("it's may be 50")
	}

	// initializer syntax, tagged
	switch i := 2 + 3; i {
	case 1:
		// single match
		fmt.Println("it's 1")
	case 2, 3, 4:
		// multiple match
		fmt.Println("it's 2, 3 or 4")
	default:
		// when everything fails
		fmt.Println("it's may be 5")
	}

	// tagless
	switch {
	case guess == 10:
		// single match
		fmt.Println("it's 10")
	case guess == 20 || guess == 30:
		// multiple match
		fmt.Println("it's 20 or 30")
		fallthrough
	default:
		// when everything fails
		fmt.Println("it's may be 50")
	}

	// type check with switch
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case float64:
		fmt.Println("i is float")
	case bool:
		fmt.Println("i is bool")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is unknown")
	}
}
