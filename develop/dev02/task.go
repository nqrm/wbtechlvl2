package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
«Задача на распаковку»
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны.

Например:

"a4bc2d5e" => "aaaabccddddde"

"abcd" => "abcd"

"45" => "" (некорректная строка)

"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.

Например:

qwe\4\5 => qwe45 (*)

qwe\45 => qwe44444 (*)

qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку.

Написать unit-тесты.
*/

func main() {
	result, err := UnpackingString(`\`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func UnpackingString(str string) (string, error) {
	strRunes := []rune(str)
	var result strings.Builder
	tmpChar := ""
	var isEscape bool
	for i := 0; i < len(strRunes); i++ {
		if unicode.IsDigit(strRunes[i]) {
			if tmpChar == "\\" {
				tmpChar = string(strRunes[i])
				//result.WriteString(tmpChar)
				//tmpChar = ""
				continue
			}
			if tmpChar == "" && !isEscape {
				return "", errors.New(fmt.Sprintf("Error: Before the digit at index %d, there must be a character (not digit).", i))
			}
			repeater, err := strconv.Atoi(string(strRunes[i]))
			if err != nil {
				return "", err
			}
			if tmpChar == "" && isEscape {
				repeater--
				tmpChar = "\\"
				isEscape = false
			}
			for j := 0; j < repeater; j++ {
				result.WriteString(tmpChar)
			}
			tmpChar = ""
		} else if strRunes[i] == '\\' {
			if tmpChar == "\\" {
				isEscape = true
				result.WriteString(tmpChar)
				tmpChar = ""
				continue
			}
			isEscape = false
			result.WriteString(tmpChar)
			tmpChar = "\\"
		} else {
			if tmpChar == "\\" {
				return "", errors.New(fmt.Sprintf("Error: Unable to escape character at index %d (%s)", i, string(strRunes[i])))
			}
			if len(tmpChar) != 0 {
				result.WriteString(tmpChar)
			}
			tmpChar = string(strRunes[i])
		}
	}
	if tmpChar == "\\" {
		return "", errors.New(fmt.Sprintf("Error: Аfter the backslash there must be either a digit or a backslash"))
	}

	result.WriteString(tmpChar)

	return result.String(), nil
}
