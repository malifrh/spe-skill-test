package main

import (
	"fmt"
	"strconv"
)

func main() {

	result := narcissistic_number(153)
	result2 := narcissistic_number(111)

	s1 := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	s2 := []int{160, 3, 1719, 19, 11, 13, -21}
	result3 := findOutlier(s1)
	result4 := findOutlier(s2)

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

}

func narcissistic_number(number int) bool {

	var tempNumber, remainder int
	var result int = 0

	tempNumber = number

	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder
		tempNumber /= 10
		if tempNumber == 0 {
			break
		}
	}

	return result == number
}

func findOutlier(arrInt []int) []string {
	var oddArr []string
	var evenArr []string

	var res1 []string
	var res2 []string

	for i := 0; i < len(arrInt); i++ {
		if arrInt[i]%2 == 0 {
			res1 = append(evenArr, strconv.Itoa(arrInt[i]))
		} else {
			res2 = append(oddArr, strconv.Itoa(arrInt[i]))
		}
	}

	if len(evenArr) == 1 {
		return res1
	} else {
		return res2
	}

	// return res
}
