package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("____Калькулятор индекса массы тела____")
	for {
		var userHeight, userKg = getUserInput()
		IMT := calculateIMT(userKg, userHeight)
		outputResult(IMT)
		isRepeat := isRepeatCheck()
		if !isRepeat {
			break
		}
	}
}

func outputResult(imt float64) {
	res := fmt.Sprintf("Ваш индекс массы тела: %.0f.", imt)
	fmt.Print(res)
	if imt < 16 {
		fmt.Println(" У вас недостаток веса.")
	}
}

func calculateIMT(kg float64, height float64) float64 {
	const IMTPower = 2
	return kg / math.Pow(height/100, IMTPower)
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите ваш рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userKg)

	return userHeight, userKg
}

func isRepeatCheck() bool {
	var answer string
	fmt.Println("Повтрорить расчет (y/n)?")
	fmt.Scan(&answer)
	if answer == "y" || answer == "Y" {
		return true
	}
	return false
}
