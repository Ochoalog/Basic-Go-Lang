package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays")

	var array1 [5]int

	fmt.Println(array1)

	array2 := [2]string{"oi", "Vitor"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(array3)

	fmt.Println("Slices: ")

	slice := []int{10, 20, 30, 40}

	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

	fmt.Println("Array internos")

	sliceMake := make([]float32, 10, 15)

	fmt.Println(sliceMake)
}
