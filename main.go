package main

import "fmt"

func basicTask()  {
	// string
	// int
	//rune
	// [capacity]type ---> 0, 1, 2, 3, 4 .... n
	//
	var array [4]int
	array[0] = 111
	array[1] = 222
	array[2] = 666
	array[3] = 999
	fmt.Println(array)
	fmt.Println("len of array =", len(array))

	/// n = 16 -----> 1 2 3 4 5 . .. 16
	for i := 0; i < len(array); i++ {
		//array[i] = 2
	}
	fmt.Println(array)

	for index, _ := range array {
		array[index]++
		//fmt.Printf("array[%d] = %d\n", index, value)
	}
	fmt.Println(array)
}

func firstTask()  {
	// [5] { a, b, c, d, e} ---- среднее арифметическое
	//Подсказка
	//TODO: ПРОЙТИСЬ ПО МАССИВУ СУММИРУЯ ВСЕ ЭЛЕМЕНТЫ В НОВУЮ ПЕРЕМЕННУЮ И ПОДЕЛИТЬ НА LEN
	var scores = [5]int{4, 2, 4, 3, 5}
	sum := 0
	for _, score := range scores{
		sum += score
	}
	var middle float64 = float64(sum) / float64(len(scores))
	fmt.Println(middle)
}
func secondTask()  {
	//////////////////////////         0          1     2     3     4    5
	var telephone = [6]string{"samsung galaxy", "sm", "sw", "sq", "se", "ip"}
	//                       {"samsung galaxy", "sam", "sw", "sam", "se", "sam"}
	for index, _ := range telephone {
		if index % 2 != 0{
			telephone[index] = "sam-gal"
		}
	}
	fmt.Println(telephone)
}

func slice()  {
	var x []string
	fmt.Println(x)
	//make
	y := make([]int, 5)
	fmt.Println(y)
	//
	z := make([]int, 5, 10)
	fmt.Println(z)

	arr := [5]float64{1, 2, 3, 4, 5}
	newArr := arr[0:4]
	newArr = arr[1:]
	newArr = arr[:]
	fmt.Println(newArr)

	// 1 2 3 4 --- > append (6) ---- > 1 2 3 4 6
	numbers := []float64{1, 2, 3, 4}
	numbers = append(numbers, 2)
	newArr = append(newArr, 3)
	fmt.Println(numbers)
	fmt.Println(newArr)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copiedQ := copy(slice2, slice1)
	fmt.Println("copied =", copiedQ)
	fmt.Println(slice1, "\n", slice2)
	slice2 = append(slice2, 2)
	fmt.Println(slice1, "\n", slice2)
}

func main() {

}