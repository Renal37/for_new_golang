package main

import (
	"fmt"
)

type Student struct {
	name  string
	grade []int
}

func (s *Student) AddGrade(grade int) {
	s.grade = append(s.grade, grade)

}

func (s Student) GetAverage() float64 {
	var aver float64
	len := len(s.grade)
	if len == 0 {
		return 0
	} else {
		for _, el := range s.grade {
			aver += float64(el)
		}
	}
	return aver / float64(len)
}

func IsExcellentStudent(i float64) bool {
	if i >= 4.5 {
		return true
	} else {
		return false
	}

}

func main() {
	var name string
	var count int
	fmt.Println("Введите имя студента:")
	fmt.Scan(&name)
	student := Student{name: name, grade: []int{}}
	fmt.Println("Сколько оценок?")
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var grade int
		fmt.Printf("Введите оценку %d: ", i+1)
		fmt.Scan(&grade)
		student.AddGrade(grade)
	}
	if IsExcellentStudent(student.GetAverage()) == true {
		fmt.Println("Средне бал студента выше или равно 4.5")
	} else {
		fmt.Println("Средне бал студента ниже 4.5")
	}
	fmt.Println("Средне арифметическое число:", student.GetAverage())

}
