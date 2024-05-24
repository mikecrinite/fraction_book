package service

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/mikecrinite/decimal_book/model"
)

func TextToDecimalFloat(str string) (float64, error) {
	result := 0.0

	curr := 0.0
	var err error
	for i, c := range str {
		fmt.Println(c)
		if runeIsWhitespace(c) {
			curr, _ = strconv.ParseFloat(model.GetDecimalValueForCharacter(' '), 64)
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

func TextToDecimalString(str string) string {
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

func TextToFractionRat(str string) big.Rat {
	float_val, err := strconv.ParseFloat(TextToDecimalString(str), 64)
	if err != nil {
		fmt.Println(err)
	}
	
	return DecimalFloatToFractionRat(float_val)
}

func DecimalFloatToFractionRat(decimal float64) big.Rat {
	rat := new(big.Rat)
	return *rat.SetFloat64(decimal)
}

func FractionRatToTextString(rat big.Rat) string {
	tempstr := convertRatToDecimal(&rat)

	return DecimalStringToText(tempstr)
}

func DecimalFloatToText(dec float64) string {
	panic("not implemented")
}

func DecimalStringToText(dec_str string) string {
	// This is inefficient but we can just keep converting to a string, splitting at the decimal, and grabbing the first 3 digits
	// I'm sure there's a way to convert the entire float to a string with no truncation

	// Format the remaining value as a string
	//tempstr := strconv.FormatFloat(float_value, 'f', 99, 64)
	// Split the string at the decimal point
	split := strings.Split(dec_str, ".")
	
	i := 0
	j := 3

	res_str := ""
	for j <= len(split[1]) {
		// Grab only the first three digits of the string
		encoded_char := split[1][i:j]
		// decode char
		decoded_char := string(model.GetCharacterForDecimalValue(encoded_char))
		// Append decoded_char to res_str
		res_str += decoded_char

		i = i + 3
		j = j + 3
	}
	
	return res_str
}

func runeIsWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

/////// CHATGPT-GENERATED CODE BELOW //////////

// convertRatToDecimal converts a big.Rat to its exact decimal representation.
func convertRatToDecimal(r *big.Rat) string {
	// Get numerator and denominator
	num := r.Num()
	den := r.Denom()

	// Initialize result string
	result := &strings.Builder{}

	// // Work with positive numbers
	// num = new(big.Int).Abs(num)
	// den = new(big.Int).Abs(den)

	// Integer part
	integerPart := new(big.Int).Div(num, den)
	result.WriteString(integerPart.String())

	// Remainder part
	remainder := new(big.Int).Mod(num, den)
	if remainder.Sign() == 0 {
		// No fractional part
		return result.String()
	}

	// Add the decimal point
	result.WriteString(".")

	// Map to store seen remainders to detect cycles
	seenRemainders := make(map[string]int)
	decimalPart := &strings.Builder{}
	position := 0

	// Perform long division
	for remainder.Sign() != 0 {
		// Check if this remainder has been seen before
		remainderStr := remainder.String()
		if pos, ok := seenRemainders[remainderStr]; ok {
			// Cycle detected
			decimalPartStr := decimalPart.String()
			decimalPartStr = decimalPartStr[:pos] + "(" + decimalPartStr[pos:] + ")"
			result.WriteString(decimalPartStr)
			return result.String()
		}

		// Remember the position of this remainder
		seenRemainders[remainderStr] = position

		// Multiply remainder by 10
		remainder.Mul(remainder, big.NewInt(10))

		// Append the integer part of the new remainder/den
		integerPart = new(big.Int).Div(remainder, den)
		decimalPart.WriteString(integerPart.String())

		// Update remainder
		remainder.Mod(remainder, den)
		position++
	}

	// No cycle detected
	result.WriteString(decimalPart.String())
	return result.String()
}

/* 
	I asked ChatGPT 4.0 to write me a function that converts a decimal representation of a string
	into a big.Rat because I had been trying to figure out how to use built-ins to do the conversion but I kept
	running into the issue where inevitably it'd convert to a float32 (or float64 or whatever big numeric type)
	and I'd lose the ability to have infinity precision because those types always end up rounding at the levels that
	we need for this program (we're going to have REALLY LONG DECIMALS). 

	The result is both the stupidest and simultaneously most intelligent solution I could have thought of. 
	It simply turns the decimal string into a String representation of an integer by multiplying it by 10 a bunch of times
	and uses that as the numerator, and then takes however many powers of 10 that required and uses that as the denominator.
	It's so stupidly simple and perfectly precise. 

	It will of course take up a ton of memory for huge inputs, so we'll eventually need to figure out how to simplify the fraction. 
	Or maybe it can't be simplified and that's the beauty of it. I've asked ChatGPT to simplify it and it was unable to.
	
	In this case, I guess that pretty much defeats the purpose of the script because it will be making the input text longer rather than shorter.

*/
func DecimalStringToFractionRat(decimal string) (*big.Rat, error) {
	// Split the string into integer and fractional parts
	parts := strings.SplitN(decimal, ".", 2)
	integerPart := parts[0]
	var fractionalPart string
	if len(parts) > 1 {
		fractionalPart = parts[1]
	}

	// Convert integer part to big integer
	integer := new(big.Int)
	if _, ok := integer.SetString(integerPart, 10); !ok {
		return nil, fmt.Errorf("invalid integer part: %s", integerPart)
	}

	// Convert fractional part to big integer if it exists
	var fractional *big.Int
	var denominator *big.Int
	if fractionalPart != "" {
		fractional = new(big.Int)
		if _, ok := fractional.SetString(fractionalPart, 10); !ok {
			return nil, fmt.Errorf("invalid fractional part: %s", fractionalPart)
		}

		// Calculate the denominator as 10^len(fractionalPart)
		denominator = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(len(fractionalPart))), nil)
	} else {
		// If there's no fractional part, the denominator is 1
		fractional = big.NewInt(0)
		denominator = big.NewInt(1)
	}

	// Calculate the numerator as integer * denominator + fractional
	numerator := new(big.Int).Mul(integer, denominator)
	numerator.Add(numerator, fractional)

	// Create the resulting big.Rat
	result := new(big.Rat).SetFrac(numerator, denominator)
	return result, nil
}

// func main() {
// 	// Example usage
// 	r := new(big.Rat)
// 	r.SetString("22/7")
// 	fmt.Println(convertRatToDecimal(r)) // "3.(142857)"
// }
