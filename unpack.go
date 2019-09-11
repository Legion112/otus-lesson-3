package unpack

import (
	"strconv"
	"strings"
	"unicode"
)

const escape = '\\'

// Parser function for unpacked strings
func Parser(s string) string {
	var unpack string
	var flagEscape bool
	data := strings.Split(s, "")
	for _, value := range data {
		if value == `\` && !flagEscape { // Check escape symbol
			flagEscape = true
		} else if flagEscape { // Next symbol is always string
			unpack, flagEscape = unpack+value, false
		} else if number, err := strconv.Atoi(value); err == nil && len(unpack) > 0 { // Check string symbol to convert int
			if number == 0 { // check zero int
				unpack = string(unpack[0 : len(unpack)-1])
			} else {
				w := string(unpack[len(unpack)-1]) // Get last string symbol
				for a := 0; a < number-1; a++ {    // Add unpack string
					unpack = unpack + w
				}
			}
		} else if err != nil {
			unpack = unpack + value // Add string symbol to result
		}
	}
	return unpack
}

// Parser2 function for unpacked strings (version 2)
func Parser2(s string) string {
	var unpack []string
	var flagEscape bool
	for _, value := range strings.Split(s, "") {
		switch {
		case value == `\` && !flagEscape:
			flagEscape = true
		case flagEscape:
			unpack, flagEscape = append(unpack, value), false
		case isNumber(value) && len(unpack) > 0:
			number, _ := strconv.Atoi(value)
			if number == 0 { // check zero int
				unpack = unpack[:len(unpack)-1]
			} else {
				w := unpack[len(unpack)-1]      // Get last string symbol
				for a := 0; a < number-1; a++ { // Add unpack string
					unpack = append(unpack, w)
				}
			}
		case !isNumber(value):
			unpack = append(unpack, value)
		}
	}
	return strings.Join(unpack, "")
}

func isNumber(s string) bool {
	var number bool
	for _, value := range s {
		number = unicode.IsNumber(value)
	}
	return number
}
