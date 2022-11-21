package main

import (
	"testing"

	"github.com/CodingSquire/go-courses/lection_2/hw/converter/converter"
)

func TestIntegerToRoman(t *testing.T) {
	cases := []struct {
		integer int
		roman   string
	}{
		{integer: 3, roman: "III"},
		{integer: 4, roman: "IV"},
		{integer: 9, roman: "IX"},
		{integer: 58, roman: "LVIII"},
		{integer: 1994, roman: "MCMXCIV"},
		{integer: 123, roman: "CXXIII"},
		{integer: 120, roman: "CXX"},
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	converterInterface := converter.NewConverter(values, symbols)

	for _, test := range cases {

		actualRoman, _ := converterInterface.IntToRoman(test.integer)

		if actualRoman != test.roman {
			t.Error("Integer:", test.integer,
				"Expected Roman value:", test.roman,
				"Actual Roman value:", actualRoman)
		}
	}
}
