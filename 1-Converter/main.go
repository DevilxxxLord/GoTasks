package main

import "fmt"

const USDEUR = 0.87
const USDRUB = 84.2

func main() {
	EURRUB := USDRUB / USDEUR
	fmt.Println(EURRUB)
}
