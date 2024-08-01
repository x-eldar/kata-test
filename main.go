package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func is_arabic(target string) bool {
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	return contains(nums, target)
}

func is_roman(target string) bool {
	nums := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	return contains(nums, target)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func calculate(num_1 int, operation string, num_2 int) int {
	switch operation {
	case "+":
		return num_1 + num_2
	case "-":
		return num_1 - num_2
	case "*":
		return num_1 * num_2
	case "/":
		if num_2 == 0 {
			panic("Деление на ноль")
		}
		return num_1 / num_2
	default:
		panic("Неверная операция")
	}
}

func romanToInt(roman string) int {
	switch roman {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		panic("Неверное римское число")
	}
}

func arabicToInt(arabic string) int {
	switch arabic {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "10":
		return 10
	default:
		panic("Неверное арабское число")
	}
}

func intToRoman(num int) string {
	romanMap := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	var keys []int
	for key := range romanMap {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	var roman string
	for i := len(keys) - 1; i >= 0; i-- {
		key := keys[i]
		symbol := romanMap[key]
		for num >= key {
			roman += symbol
			num -= key
		}
	}

	return roman
}

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")

	if len(parts) != 3 {
		panic("Некорректный ввод")
	}

	a := parts[0]
	operation := parts[1]
	b := parts[2]

	if is_arabic(a) && is_arabic(b) {
		result := calculate(arabicToInt(a), operation, arabicToInt(b))
		fmt.Println(result)
	} else if is_roman(a) && is_roman(b) {
		result := calculate(romanToInt(a), operation, romanToInt(b))
		if result < 1 {
			panic("Отрицательное число или ноль в римской системе счисления недопустимы")
		} else {
			romanResult := intToRoman(result)
			fmt.Println(romanResult)
		}
	} else {
		panic("Неверная операция")
	}
}
