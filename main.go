package main

import (
	"fmt"

	//"github.com/mikecrinite/decimal_book/model"
	"github.com/mikecrinite/decimal_book/service"
)

func main() {
	doStuff()
}

func doStuff() {
	str := "This is a test of the emergency alert system."

	// res_dec, err := service.TranslateTextToDecimal(str)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	frac := service.TranslateTextToFraction(str)

	//res_str := service.TranslateTextToDecimalString(str)
	//fmt.Println(res_dec)
	//fmt.Println(res_str)

	// convert to fraction as that is the end goal
	//frac := service.DecimalToFraction(res_dec)
	res_str := service.TranslateFractionToText(frac)

	fmt.Println("======= Fraction Value: ==========")
	fmt.Println(frac.Num().String() + "/" + frac.Denom().String())
	fmt.Println("======== String Value: ===========")
	fmt.Println(res_str)
}
