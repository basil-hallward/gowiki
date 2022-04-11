package main

import "fmt"

type testInt func(int) bool // 関数の型を変数として定義する

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// ふたつの関数の処理を、ひとつの関数内にまとめて記述することができる
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			if f(value) {
				result = append(result, value)
			}
		}
	}
	return result
}

func main() {
	slice := []int {1, 2, 3, 4, 5, 7}

	odd := filter(slice,  isOdd)
	fmt.Println("Odd elements of slice are: ", odd)

	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)
}