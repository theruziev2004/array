package main

import "fmt"

func main() {
	//int
	//string
	//rune
	//[вместимость]type ---> 0, 1, 2, 3 .... n

	var array [4]int
	array[0] = 111
	array[1] = 333
	array[2] = 666
	array[3] = 999

	fmt.Println("len of array = ", len(array))

/*	i := 0
	for {
		if i >= len(array){
			break
		}
		array[i] = 1
		i++
	}*/

	for i := 0; i < len(array); i++ {
		//array[i] = 11
	}
	fmt.Println(array)

	for index, value := range array{
		//array[index]++
		fmt.Printf("array[%d] = %d\n", index, value)
	}
	fmt.Println(array)
	//                            0               1     2     3     4       5
	var telephone = [6]string{"samsung galaxy", "sm", "sw", "sq", "se", "iphone"}
	for i := 0; i < len(telephone); i++ {
		if i % 2 != 0 {
			telephone[i] = "samsung galaxy"
		}
	}
	fmt.Println(telephone)
}


//// [5]int {a, b, c , d, e}----- среднее арифметическое
