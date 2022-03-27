package main

import "fmt"

func main() {
	var (
		aInt = 55
		bInt = 23
	)
	fmt.Println(aInt + bInt)
	fmt.Println(aInt - bInt)
	fmt.Println(aInt * bInt)
	fmt.Println(aInt / bInt) // integer division will result in integer output
	fmt.Println(aInt % bInt) // only available for Int

	var (
		aFloat = 3.8
		bFloat = 1.33
	)
	fmt.Println(aFloat + bFloat)
	fmt.Println(aFloat - bFloat)
	fmt.Println(aFloat * bFloat)
	fmt.Println(aFloat / bFloat)

	// bit operators, only integer
	fmt.Println(aInt & bInt)  // and
	fmt.Println(aInt | bInt)  // or
	fmt.Println(aInt ^ bInt)  // xor
	fmt.Println(aInt &^ bInt) // and not

	// bit shifting operator, only integer
	a := 8
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0
}
