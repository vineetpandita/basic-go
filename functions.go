package main

import "fmt"

func main() {


	variadicfunction(1,2,3,4,5,6,6)

	data := []int {13,45,66,66,77,76,654}

	//Argumnts getting converted to variadic

	variadicfunction(data...)  // sending a slice as a seperate values

}


//variadic

//Parameter dots are at front!


func variadicfunction (i ...int){

	fmt.Println(i) // i is a slice

	for _, val := range i {
		fmt.Println(val)
	}


}

