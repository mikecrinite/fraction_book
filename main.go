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

	frac := service.TranslateTextToFraction(str)
	
	res_str := service.TranslateFractionToText(frac)

	fmt.Println("======= Fraction Value: ==========")
	fmt.Println(frac.Num().String() + "/" + frac.Denom().String())
	fmt.Println("======== String Value: ===========")
	fmt.Println(res_str)
}
