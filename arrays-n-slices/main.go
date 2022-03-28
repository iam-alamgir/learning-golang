package main

import "fmt"

func main() {
	// arrays	:	variable type, not reference, fixed size
	var a [3]int // defines a fixed size array
	fmt.Printf("%v, %T\n", a, a)

	b := [...]int{1, 2, 3, 4} // defines an array of exact size of provided values dynamically
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", b[1], b[1])
	fmt.Printf("size: %v\n", len(b))

	c := a // creates a separate copy
	fmt.Printf("%v, %T\n", c, c)
	c[1] = 3
	fmt.Printf("%v, %T\n", c, c)

	d := &b // points to the same array
	fmt.Printf("%v, %T\n", d, d)

	// slices	:	reference type, resizable
	l := make([]int, 3)
	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("size: %v\n", len(l))
	fmt.Printf("capacity: %v\n", cap(l))

	m := make([]int, 3, 20)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("size: %v\n", len(m))
	fmt.Printf("capacity: %v\n", cap(m))

	x := []int{1, 2, 3}
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("size: %v\n", len(x))
	fmt.Printf("capacity: %v\n", cap(x))

	y := x // points to the same slice
	fmt.Printf("%v, %T\n", y, y)

	// slices from slice or array
	z := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := z[:] // creates slice of all values
	fmt.Printf("%v, %T\n", p, p)
	q := z[3:] // create slice of values from index 3 (including) till end
	fmt.Printf("%v, %T\n", q, q)
	r := z[:6] // create slice of values from start till index 6 (excluding)
	fmt.Printf("%v, %T\n", r, r)
	s := z[3:7] // create slice of values from index 3 (including) till index 7 (excluding)
	fmt.Printf("%v, %T\n", s, s)

	t := []int{1, 2, 3, 4, 5}
	u := []int{6, 7, 8, 9, 10}

	j := append(t, u...)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v\n", j[1:])                   // trim from start
	fmt.Printf("%v\n", j[:len(j)-1])            // trim from end
	fmt.Printf("%v\n", append(j[:2], j[3:]...)) // remove middle, causes issue as passed by reference
	fmt.Printf("%v\n", j)                       // repeat issue due to previous statement, does not occur for removing more than one element
}
