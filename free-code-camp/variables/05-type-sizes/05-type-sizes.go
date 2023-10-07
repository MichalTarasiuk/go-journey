package main

import "fmt"

func main() {
	accountAge := 2.6
	accountAgeInt := int(accountAge)

	fmt.Println("var accountAgeInt =", accountAgeInt)

	temperatureInt := 88
	temperatureFloat := float64(temperatureInt)

	fmt.Println("var temperatureFloat =", temperatureFloat)
}
