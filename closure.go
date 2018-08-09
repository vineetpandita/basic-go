package main

import (
	"fmt"
)

func powerOf (power int) func(int) int {

	return func(a int) int {


		return a^power; // need to fix this this is negation.
    }
	
}


func lookCool( a int, callback func(int) int) int{

	return callback(a)

}


func filter(numbers []int, callback func(int) bool) []int{

	xs := []int{};

	for _, n := range numbers {
		if(callback(n)){
			xs = append(xs,n)
		}
	}

	return xs

}



func main() {

	square := powerOf(2);
	cube   := powerOf(3);


	fmt.Println("Cube of 10 is ", cube(10))
	fmt.Println("Square of 10 is ", square(10))


	fmt.Println("Cool Cube of 10 is ", lookCool(10, cube))
	defer fmt.Println("Cool Square of 10 is ", lookCool(10, square))


	xs4 := filter([]int{1,2,3,4,5,5}, func(n int) bool {
	   return n>4;
	})

	defer fmt.Println(xs4)


	xs2 := filter([]int{1,2,3,4,5,5}, func(n int) bool {
		return n>2;
	})

	fmt.Println(xs2)



	// Self executing function

	func () {

		fmt.Println(" Executed Self")
	}()


	//io.Reader()

}

