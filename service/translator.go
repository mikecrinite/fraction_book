package service

import (
	"strconv"
	"unicode"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/mikecrinite/decimal_book/model"
)

func TranslateTextToDecimal(str string) (float64, error) {
	result := 0.0

	curr := 0.0
	var err error
	for i, c := range str {
		fmt.Println(c)
		if runeIsWhitespace(c) {
			curr, err = strconv.ParseFloat(model.GetDecimalValueForCharacter(' '), 64)
		} else {
			curr, err = strconv.ParseFloat(model.GetDecimalValueForCharacter(c), 64)
			if err != nil {
				fmt.Println(err)
			}
		}

		offset := math.Pow(1000, float64(-1 * (i + 1)))
		result += curr * offset

		//fmt.Println(fmt.Sprintf("%s : %d", string(c), result))
	}

	return result, nil
}

func TranslateTextToDecimalString(str string) string {
	result := "0."

	for _, c := range str {
		if runeIsWhitespace(c) {
			result += model.GetDecimalValueForCharacter(' ')
		} else {
			result += model.GetDecimalValueForCharacter(c)
		}
	}

	return result
}

func TranslateTextToFraction(str string) big.Rat {
	float_val, err := strconv.ParseFloat(TranslateTextToDecimalString(str), 64)
	if err != nil {
		fmt.Println(err)
	}
	
	return DecimalToFraction(float_val)
}

func DecimalToFraction(decimal float64) big.Rat {
	rat := new(big.Rat)
	return *rat.SetFloat64(decimal)
}

func TranslateFractionToText(rat big.Rat) string {
	float_value, exact := rat.Float64()

	if !exact {
		fmt.Println("The fraction to float conversion was not exact. The text may not have been encoded properly")
	}


	fmt.Println(float_value)
	curr := float_value

	var tempstr string
	res_str := ""
	//var err error

	// This is inefficient but we can just keep converting to a string, splitting at the decimal, and grabbing the first 3 digits
	// I'm sure there's a way to convert the entire float to a string with no truncation
	i := 1.0
	for curr > 0.0 {
		//TODO: FIGURE OUT HOW NOT TO TRUNCATE THE FLOAT VALUE!!!

		// Format the remaining value as a string
		tempstr = fmt.Sprintf("%f", curr)
		// Split the string at the decimal point
		split := strings.Split(tempstr, ".")
		// Grab only the first three digits of the string
		encoded_char := split[1][0:3]
		// decode char
		decoded_char := string(model.GetCharacterForDecimalValue(encoded_char))
		// Append decoded_char to res_str
		res_str += decoded_char

		fmt.Println("-------------------------------")
		fmt.Println(tempstr)
		//fmt.Println(encoded_char)
		fmt.Println(decoded_char)
		fmt.Println(res_str)
		curr = math.Pow(1000, i) * float_value
		i++
		//fmt.Println(encoded_char)
	}

	return ""
}

func TranslateDecimalToText(dec float64) string {
	return ""
}

func runeIsWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}