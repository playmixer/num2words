package num2words

import (
	"strings"
)

const (
	THOUSAND = 1000
	MILLION  = 1000000
	MILLIARD = 1000000000
)

var _smallNumbers = []string{
	"ноль", "один", "два", "три", "четыре",
	"пять", "шесть", "семь", "восемь", "девять",
	"десять", "одинадцать", "двенадцать", "тринадцать", "четырнадцать",
	"пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать",
}

var _tens = []string{
	"", "", "двадцать", "тридцать", "сорок", "пятьдесят",
	"шестьдесят", "семьдесят", "восемьдесят", "девяносто",
}

var _hundreds = []string{
	"", "сто", "двести", "триста", "четыреста", "пятьсот",
	"шестьсот", "семьсот", "восемьсот", "девятьсот",
}

var _scaleNumbers = [][]string{
	{"", "", ""}, {"тысяча", "тысячи", "тысячь"}, {"миллион", "миллиона", "миллионов"}, {"миллиард", "миллиарда", "миллиардов"},
}

func Convert(num int) string {
	minus := ""
	_num := num
	if num < 0 {
		minus = "минус "
		_num = num * -1
	}
	return strings.TrimSpace(minus + convert(_num))
}

func convert(num int) string {
	smal := num % 100

	res := _smallNumbers[smal%10]
	if smal < 20 && smal > 0 {
		res = _smallNumbers[smal]
	}
	if num > 0 && smal%10 == 0 && smal != 10 {
		res = ""
	}

	return strings.TrimSpace(tens(num) + " " + res)
}

func tens(num int) string {
	_ten := num % 100 / 10

	return strings.TrimSpace(hundred(num) + " " + _tens[_ten])
}

func hundred(num int) string {
	_hunderd := num % 1000 / 100
	return strings.TrimSpace(scale(num) + " " + _hundreds[_hunderd])
}

func scale(num int) string {
	_milliard := num % 1000000000000 / MILLIARD
	if _milliard > 0 {
		return strings.TrimSpace(convert(_milliard) + " " + NumbericDeclensio(_milliard, _scaleNumbers[3]) + " " + scale(num%MILLIARD))
	}

	_million := num % MILLIARD / MILLION
	if _million > 0 {
		return strings.TrimSpace(convert(_million) + " " + NumbericDeclensio(_million, _scaleNumbers[2]) + " " + scale(num%MILLION))
	}

	_thousand := num % MILLION / THOUSAND
	if _thousand > 0 {
		return strings.TrimSpace(groupDeclensio(_thousand) + " " + NumbericDeclensio(_thousand, _scaleNumbers[1]))
	}

	return ""
}

func groupDeclensio(num int) string {

	_c := ""
	switch num {
	case 0:
	case 1:
		_c = "одна"
	case 2:
		_c = "две"
	default:
		_c = convert(num)
	}
	return _c
}

func NumbericDeclensio(number int, words []string) string {
	remainder10 := number % 10
	remainder100 := number % 100

	switch {
	case remainder100 >= 11 && remainder100 <= 20:
		return words[2]
	case remainder10 == 1:
		return words[0]
	case remainder10 >= 2 && remainder10 <= 4:
		return words[1]
	default:
		return words[2]
	}
}
