package main

import (
	"fmt"
)

func konvertTemp(temp int,v int) float64 {
	if v == 1 {
		return (float64(temp) - 32) * (5.0 / 9.0)
	} else if v == 2 {
		return (float64(temp) * (9.0 / 5.0)) + 32
	} else {
		return 0
	}
}

func main() {
	var variant int
	var temp int 
	fmt.Printf("Выберите вариант конвертации \nФаренгейтом в Цельсии - 1 \nЦельсии в Фаренгейтом - 2\n")
	fmt.Scan(&variant)
	fmt.Println("Введите температуру")
	fmt.Scan(&temp)
	fmt.Println("Конвертированная температура",konvertTemp(temp,variant))
}
