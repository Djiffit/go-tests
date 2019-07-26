package roman

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var numerals = []RomanNumeral{
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

// ConvertToRoman numeral a number
func ConvertToRoman(num int) string {
	var result strings.Builder

	for _, pair := range numerals {
		for num >= pair.Value {
			result.WriteString(pair.Symbol)
			num -= pair.Value
		}
	}

	return result.String()
}

func ConvertToNumber(roman string) int {
	result := 0

	for _, pair := range numerals {
		for isStart(pair.Symbol, roman) {
			result += pair.Value
			roman = roman[len(pair.Symbol):]
		}
	}

	return result
}

func isStart(symbol, roman string) bool {
	if len(symbol) > len(roman) {
		return false
	}

	return roman[:len(symbol)] == symbol
}
