package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("GoCalculator")
	MathOperation := strings.Fields(Scaner())
	switch {
	case len(MathOperation) == 1:
		fmt.Println("Ошибка ввода! Cтрока не является математической операцией")
		return
	case len(MathOperation) == 2:
		fmt.Println("Ошибка ввода! Возможно отсутствует знак математической операции")
		return
	case len(MathOperation) > 3:
		fmt.Println("Ошибка ввода! Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	}
	var operation = MathOperation[1]
	var FirstNumber, SecondNumber int
	var RimDigits bool
	switch {
	case RimDigitsToInt(MathOperation[0]) != 0 && RimDigitsToInt(MathOperation[2]) != 0:
		FirstNumber = RimDigitsToInt(MathOperation[0])
		SecondNumber = RimDigitsToInt(MathOperation[2])
		RimDigits = true
		if (FirstNumber <= SecondNumber) && (operation == "-") {
			fmt.Println("Ошибка, в Римской системе нет нуля и отрицательных чисел!")
			return
		}
	case (RimDigitsToInt(MathOperation[0]) == 0 && RimDigitsToInt(MathOperation[2]) != 0) ||
		(RimDigitsToInt(MathOperation[0]) != 0 && RimDigitsToInt(MathOperation[2]) == 0):
		fmt.Println("Ошибка, используются разные системы счисления")
		return
	case RimDigitsToInt(MathOperation[0]) == 0 && RimDigitsToInt(MathOperation[2]) == 0:
		FirstNumber, _ = strconv.Atoi(MathOperation[0])
		SecondNumber, _ = strconv.Atoi(MathOperation[2])
	}

	if FirstNumber > 10 || FirstNumber < 1 || SecondNumber > 10 || SecondNumber < 1 {
		fmt.Println("Ошибка, операции производятся с целыми Римскими или Арабскими числами от 1 до 10 включительно!")
		return
	}
	result, err := Calc(FirstNumber, operation, SecondNumber)
	if err == nil {
		if RimDigits {
			fmt.Println(RimToArab(result))
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println(err)
	}
}

func Scaner() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return strings.ToUpper(in.Text())
}

func RimToArab(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	rim := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			rim += conversion.digit
			number -= conversion.value
		}
	}
	return rim
}

func RimDigitsToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func Calc(a int, action string, b int) (int, error) {
	switch action {
	case "*":
		return a * b, nil
	case "-":
		return a - b, nil
	case "+":
		return a + b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("ошибка ввода - недопустимая операция")
	}
}
