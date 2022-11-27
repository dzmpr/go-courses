package converter

import "strings"

type converterImpl struct {
	symbols []string
	values  []int
}

type Converter interface {
	IntToRoman(input int) string
}

// IntToRoman Converts given int number to string with roman representation
func (c *converterImpl) IntToRoman(num int) string {
	var res strings.Builder

	for i, val := range c.values {
		if num >= val {
			var count int = num / val
			res.WriteString(strings.Repeat(c.symbols[i], count))
			num -= count * val
		}
	}

	return res.String()
}

func NewConverter(values []int, symbols []string) Converter {
	return &converterImpl{
		values:  values,
		symbols: symbols,
	}
}
