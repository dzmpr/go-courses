package converter

import (
	"strings"
)

type converterImpl struct {
	symbols []string
	values  []int
}

type Converter interface {
	IntToRoman(input int) (string, error)
}

func (impl *converterImpl) IntToRoman(num int) (string, error) {
	var res strings.Builder

	for i, val := range impl.values {
		if num >= val {
			var count int = num / val
			res.WriteString(strings.Repeat(impl.symbols[i], count))
			num -= count * val
		}
	}

	return res.String(), nil
}

func NewConverter(values []int, symbols []string) Converter {
	return &converterImpl{
		values:  values,
		symbols: symbols,
	}
}

func IntToRoman(input int) {

}
