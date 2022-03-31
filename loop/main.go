package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("%v\t", i)
	}
	fmt.Println()

	// cannot use i++ / j++ for this increment for multiple values
	for i, j := 0, 0; i < 4; i, j = i+1, j+1 {
		fmt.Printf("%v, %v\n", i, j)
	}

	i := 0
	// must keep the semicolon at the start
	for ; i < 4; i++ {
		fmt.Printf("%v\t", i)
	}
	fmt.Println()

	j := 0
	// while loop
	for j < 4 {
		fmt.Printf("%v\t", j)
		j++ // must update value in the loop to avoid infinite looping
	}
	fmt.Println()

	k := 0
	// works the same as previous loop (while loop shorthand)
	for k < 4 {
		fmt.Printf("%v\t", k)
		k++ // must update value in the loop to avoid infinite looping
	}
	fmt.Println()

	l := 0
	// infinite loop
	for {
		fmt.Printf("%v\t", l)
		if l > 3 {
			break // exits the loop
		}
		l++
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skips the rest of the loop and goes to next iteration
		}
		fmt.Printf("%v\t", i)
	}
	fmt.Println()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i*j > 3 {
				break // breaks out of inner loop only
			}
			fmt.Printf("%v\t", i*j)
		}
	}
	fmt.Println()

ExitHere: // label
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i*j > 3 {
				break ExitHere // breaks out to label
			}
			fmt.Printf("%v\t", i*j)
		}
	}
	fmt.Println()

	s := []int{1, 2, 3}
	// works with arrays, slices and maps
	// k and v are not fixed
	// to ignore value, remove `v` altogether
	// to ignore key, replace `k` with `_`
	for k, v := range s {
		fmt.Printf("%v, %v\t", k, v)
	}
	fmt.Println()
}
