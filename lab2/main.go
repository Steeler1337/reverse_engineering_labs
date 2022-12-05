package main

import "fmt"

func main() {
	var array = new([2]int)
	array[0] = 5
	array[1] = 6
	//fmt.Print(array[0])

	//array = nil

	fmt.Print(array)

	var double_array = new([2][2]int)
	double_array[0][0] = 1118481
	double_array[0][1] = 2236962
	double_array[1][0] = 3355443
	double_array[1][1] = 4473924
	//fmt.Print(double_array[0][0])
	double_array = nil

	fmt.Print(double_array)

}
