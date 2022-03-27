package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		aInt    = 2354
		aFloat  = 55.12
		aString = "Hello"
		bString = " Golang!!"
	)

	// Conversion
	convI := float64(aInt)
	convJ := int(aFloat)
	fmt.Printf("%v, %T\n", convI, convI)
	fmt.Printf("%v, %T\n", convJ, convJ)

	convIntData := strconv.Itoa(aInt)
	fmt.Printf("%v, %T\n", convIntData, convIntData)

	convFloatData := fmt.Sprintf("%f", aFloat)
	fmt.Printf("%v, %T\n", convFloatData, convFloatData)

	convAStringData, err := strconv.ParseFloat(convFloatData, 32)
	convBStringData, err := strconv.ParseFloat(convFloatData, 64) // gives actual precision
	fmt.Printf("%v, %T\n", convAStringData, convAStringData)
	fmt.Printf("%v, %T\n", convBStringData, convBStringData)
	fmt.Printf("%v, %T\n", err, err)

	fmt.Printf("%v, %T\n", aString+bString, aString+bString)
}
