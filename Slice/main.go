package main

import "fmt"

func main() {
	//Array
	var a [10]int
	a[0] = 9
	a[1] = 8

	fmt.Println(a, a[0])

	var b = [2]string{"apple", "samsung"}
	fmt.Println(b)

	//Slice - mutable
	var c []int
	fmt.Println(c)

	var d = []string{
		"apple",
		"samsung",
	}
	d = append(d, "Nokia")
	d = append(d, "Sony")

	fmt.Println(d)
	fmt.Println(d[0:2])
	fmt.Println(d[:2])
	fmt.Println(d[2:])

	// fmt.Println("len:%d, cap:%d, %p \n", len(d), cap(d), d)
}
