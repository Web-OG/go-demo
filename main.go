package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("____Калькулятор индекса массы тела____")
	for {
		var userHeight, userKg, err = getUserInput()
		fmt.Printf("Произошла ошибка: %v \n", err)
		if err != nil {
			break
		}
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

func getUserInput() (float64, float64, error) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите ваш рост: ")
	_, err := fmt.Scan(&userHeight)
	if err != nil {
		return 0, 0, errors.New("Не верный ввод роста")
	}
	fmt.Print("Введите ваш вес: ")
	_, err = fmt.Scan(&userKg)
	if err != nil {
		return 0, 0, errors.New("Не верный ввод веса")
	}

	return userHeight, userKg, nil
}

func isRepeatCheck() bool {
	var answer string
	fmt.Println("Повтрорить расчет (y/n)?")
	_, err := fmt.Scan(&answer)
	if err != nil {
		return false
	}
	if answer == "y" || answer == "Y" {
		return true
	}
	return false
}
