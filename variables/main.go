package main

import (
	"fmt"
)

var i int = 34 // package level variable, must be explicit declaration, package scoped
var I int = 23 // Pascal case, public access

func main() {
	// Declaration
	// split declaration assigns default zero value in variable
	// --------------------------------------------------------------------------------
	// var i int		// split declaration, explicit
	// i = 1			// assignment
	// var i int = 1	// single line declaration and assignment, explicit
	// var i = 1		// single line declaration and assignment, inferred
	// i := 1			// single line declaration and assignment, inferred
	// var (			// block type declaration
	// 	 firstName = "John"
	// 	 lastName  = "Doe"
	// 	 age       = 33
	// )
	// --------------------------------------------------------------------------------

	// accessing global variable
	// shadowed variables can not be accessed
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", I, I)

	// integer data
	var (
		aInt8  int8  = 2
		aInt16 int16 = 23
		aInt32 int32 = 34
		aInt64 int64 = 2354
	)
	fmt.Printf("%v, %T\n", aInt8, aInt8)
	fmt.Printf("%v, %T\n", aInt16, aInt16)
	fmt.Printf("%v, %T\n", aInt32, aInt32)
	fmt.Printf("%v, %T\n", aInt64, aInt64)

	// floating point data
	var (
		aFloat32 float32 = 2.45
		aFloat64 float64 = 23.154 // defaults to float64
		bFloat64         = 3e12
		cFloat64         = 2e12
	)
	fmt.Printf("%v, %T\n", aFloat32, aFloat32)
	fmt.Printf("%v, %T\n", aFloat64, aFloat64)
	fmt.Printf("%v, %T\n", bFloat64, bFloat64)
	fmt.Printf("%v, %T\n", cFloat64, cFloat64)

	// boolean data
	aBool := true
	fmt.Printf("%v, %T\n", aBool, aBool)

	// complex data
	var aComplex64 complex64 = complex(1, 5) // result -> 1 + 5i
	var aComplex128 complex128 = 1 + 2i      // defaults to complex128
	fmt.Printf("%v, %T\n", aComplex64, aComplex64)
	fmt.Printf("%v, %T\n", aComplex128, aComplex128)
	fmt.Printf("%v, %T\n", real(aComplex128), real(aComplex128))
	fmt.Printf("%v, %T\n", imag(aComplex128), imag(aComplex128))

	// string data
	var aString string = "Hello from GoLand"
	fmt.Printf("%v, %T\n", aString, aString)       // all utf-8
	fmt.Printf("%v, %T\n", aString[2], aString[2]) // prints ascii value
	fmt.Printf("%v, %T\n", string(aString[2]), aString[2])

	bytes := []byte(aString) // byte slices
	fmt.Printf("%v, %T\n", bytes, bytes)

	// utf-32 data, alias for int32
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

	// constants, must initialize at compile-time
	const constInt = 33
	fmt.Println(constInt)
}
