// Package fizzbuzz
//
// Reference:
// https://ja.wikipedia.org/wiki/Fizz_Buzz
package fizzbuzz

import "strconv"

func Fizzbuzz(number int) string {

	if number%15 == 0 {
		return "Fizz Buzz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(number)
}
