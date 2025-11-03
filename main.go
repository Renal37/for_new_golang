package main

import (
	"fmt"
)

func sumArray(numbers []int) int {
	var sum int
	len := len(numbers)
	if len != 0 {
		for _, el := range numbers {
				sum +=el
		}
	} else {
		return sum
	}
	return sum
}

func main() {
var nums []int
	var num int
	
	fmt.Println("Введите числа (для завершения введите любой нечисловой символ):")
	
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		nums = append(nums, num)
	}
	
	fmt.Println("Исходный срез:", nums)
	fmt.Println("Сумма:", sumArray(nums))
}
