package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64
	var userKg float64
	fmt.Println("____Калькулятор индекса массы тела____")
	fmt.Print("Введите ваш рос: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userKg)
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	res := fmt.Sprintf("ваш индекс массы тела: %.0f", imt)
	fmt.Print(res)
}

func calculateIMT(kg float64, height float64) float64 {
	const IMTPower = 2
	return kg / math.Pow(height/100, IMTPower)
}
