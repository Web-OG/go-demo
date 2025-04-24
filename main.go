package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Println("____Калькулятор индекса массы тела____")
	fmt.Print("Введите ваш рос: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userKg)
	var IMT = userKg / math.Pow(userHeight/100, IMTPower)
	outputResult(IMT)
}

func outputResult(imt float64) {
	res := fmt.Sprintf("ваш индекс массы тела: %.0f", imt)
	fmt.Print(res)
}
