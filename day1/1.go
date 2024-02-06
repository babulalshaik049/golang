package main

//This is for single line comment
import (
	"fmt"
	"reflect" // This module is for identifying the type of a variable
)

func main() {
	fmt.Println("Hello world")
	/* this is for
	multi-line
	comment */
	//Below code is understanding datatypes
	a := true
	b := "string"
	c := 12
	d := 14.0
	fmt.Println(a)
	fmt.Println("a value is:", a, "and data type of variable a is ", reflect.TypeOf(a))
	fmt.Println("b value is:", b, "and data type of variable b is ", reflect.TypeOf(b))
	fmt.Println("b value is:", c, "and data type of variable b is ", reflect.TypeOf(c))
	fmt.Println("b value is:", d, "and data type of variable b is ", reflect.TypeOf(d))
	fmt.Println("-------------------------------------------------------------------------------------------------------------------")
	//Arrays are going to be of same datatype syntax: var arrayname [size]ElementDataType

	first_array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("first_array value is:", first_array, "and data type of variable first_array is ", reflect.TypeOf(first_array))
	fmt.Println("-------------------------------------------------------------------------------------------------------------------")

	//if you are not sure of size of an array, you can try like below
	second_array := [...]int{12, 13, 14}
	fmt.Println("second_array value is:", second_array, "and data type of variable second_array is ", reflect.TypeOf(second_array))
	fmt.Println(len(second_array))
	fmt.Println("-------------------------------------------------------------------------------------------------------------------")

	//using loops on arrays
	for i := 0; i < len(second_array); i++ {
		fmt.Println(second_array[i])
	}

	fmt.Println("====================================================================================================================")

}
