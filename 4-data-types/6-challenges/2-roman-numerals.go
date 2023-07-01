package main

import (
	"fmt"
)

func convertRomanNumeralStringToInt(romanNumeral string) int {
	romanNumeralMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var total int

	intVals := make([]int, len(romanNumeral)+1)

	for i, digit := range romanNumeral {
		currentSymbolValue, present := romanNumeralMap[digit]
		if present {
			intVals[i] = currentSymbolValue
		} else {
			fmt.Printf("Invalid roman numeral %s, unrecognised digit: %c\n", romanNumeral, digit)
			return 0
		}
	}

	for i := 0; i < len(romanNumeral); i++ {
		if intVals[i] < intVals[i+1] {
			total -= intVals[i]
		} else {
			total += intVals[i]
		}
	}

	return total
}

func main() {
	fmt.Println("Enter a roman numeral: ")
	var romanNumeral string
	fmt.Scanf("%s", &romanNumeral)
	fmt.Println(convertRomanNumeralStringToInt(romanNumeral))
}

// func TestConvert(t *testing.T) {
// 	symbol := "MCMXCVII"
// 	want := 1997
// 	value := convertRomanNumeralStringToInt(symbol)
// 	if want != value {
// 		t.Fatalf(`convertRomanNumeralStringToInt(%s) = %v, want match for %#q, nil`,
// 			symbol, value, want)
// 	}
// }
