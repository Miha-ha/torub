package torub

import "testing"

func TestConvert(t *testing.T) {
	tests := map[string]string{
		"123.34":    "Сто двадцать три рубля 34 коп.",
		"1234.56":   "Одна тысяча двести тридцать четыре рубля 56 коп.",
		"123456.76": "Сто двадцать три тысячи четыреста пятьдесят шесть рублей 76 коп.",
	}

	for num, rub := range tests {
		if Convert(num) != rub {
			t.Errorf("Error: %v != %v\n", num, rub)
		}
	}
}
