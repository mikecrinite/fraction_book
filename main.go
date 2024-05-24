package main

import (
	"fmt"

	"github.com/mikecrinite/decimal_book/service"
)

func main() {
	doStuff()
}

func doStuff() {
	str := "This is a test of the emergency alert system"

	// The first step is to convert the text into a String which represents the decimal we're trying to convert
	decimal_string := service.TextToDecimalString(str)
	fmt.Println("======= Decimal String: ==========")
	fmt.Println(decimal_string)
	// We can test this conversion by now converting back to the text
	result_str_one := service.DecimalStringToText(decimal_string)
	fmt.Println("======= Text From Decimal String: ==========")
	fmt.Println(result_str_one)

	// The next step would be to convert that String decimal value into a Rat, in order to represent it as a fraction
	// This decimal representation would ideally be a simplified fraction. However, what I actually discovered is that the 
	// simplest fraction would MUCH longer than the original text. So at the end of the day, we have learning nothing.
	// I guess in retrospect, if this was possible to do, someone would have already done it in order to more efficiently 
	// store text
	frac, err := service.DecimalStringToFractionRat(decimal_string)
	if err != nil {
		println(err)
	}

	// 
	res_str := service.FractionRatToTextString(*frac)

	fmt.Println("======= Fraction String: ==========")
	fmt.Println(frac.Num().String() + "/" + frac.Denom().String())
	fmt.Println("======== Result Text: ===========")
	fmt.Println(res_str)
}
