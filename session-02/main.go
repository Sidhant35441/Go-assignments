package main

import "fmt"

//converting slices to single integer
func SliceToInt(numbers []int) int {
	var i, k int
	k = 0
	for i = 0; i < len(numbers); i++ {
		k = 10*k + numbers[i]
	}
	return k
}
func Add(number1 []int, number2 []int) string {
	//reversing number2 slice
	for i, j := 0, len(number2)-1; i < j; i, j = i+1, j-1 {
		number2[i], number2[j] = number2[j], number2[i]
	}
	//addition operation
	result := fmt.Sprintf("result: %v", SliceToInt(number1)+SliceToInt(number2))
	return result
}

func main() {
	n1 := []int{1, 2, 3, 4, 5, 6}
	n2 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(Add(n1, n2))
}
