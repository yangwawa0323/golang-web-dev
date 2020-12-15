package main

import "fmt"

// slice
// map
// struct

type person struct {
	fname string
	lname string
}

func main() {
	// initial an Integer array
	xi := []int{2, 3, 4, 5, 6, 7}
	xi = append(xi)
	fmt.Println(xi)

	// use make function create a string array, with zero element
	// and set the capacity to 100.
	xs := make([]string, 0, 100)
	// append element will be the first element `xs[0]`
	xs = append(xs, "Go rocks.")
	fmt.Println("after append xs: ", xs)
	fmt.Println("xs length is: ", len(xs))
	fmt.Println("xs capacity is: ", cap(xs))
	fmt.Println("xs[0] is: ", xs[0])

	for i, v := range xi {
		fmt.Println(i, v)
	}

	// initial a map...
	y := map[string]string{
		"Todd": "chocolate",
		"Jim":  "vanilla",
	}
	y["Mark"] = "chocolate also"

	fmt.Println(y)

	for k, v := range y {
		fmt.Println(k, v)
	}

	z := &person{
		"James",
		"Bond",
	}

	// Access a struct by reference pointer or by value is not difference.
	fmt.Println("z is a pointer  of person struct: ", z)
	fmt.Println("use z pointer's attribute: ", z.fname)
	fmt.Printf("%T\n", z)
}
