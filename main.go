package main

import (
	"flag"
	"fmt"
)

// TotalCost to calc total cost, tip and cost per person
func TotalCost(Price, Tax, Tip float64, People uint) (float64, float64, float64) {
	var taxConverted float64 = Tax / 100.0
	var tipConverted float64 = (Tip / 100.0)
	var totalTip float64 = (Price * tipConverted)
	var totalTot float64 = Price + (Price * taxConverted) + totalTip
	var costPerPerson float64 = totalTot / float64(People)
	return totalTip, totalTot, costPerPerson
}

type userInput struct {
	Price  float64
	Tax    float64
	Tip    float64
	People uint
}

func main() {
	//args
	var Price = flag.Float64("Price", 100.0, "Price of goods")
	var Tax = flag.Float64("Tax", 8.5, "State food tax")
	var Tip = flag.Float64("Tip", 15.0, "Percentage to Tip")
	var People = flag.Uint("People", 3, "Number of people")
	flag.Parse()

	u := userInput{*Price, *Tax, *Tip, *People}

	var tip, total, costPerPerson = TotalCost(u.Price, u.Tax, u.Tip, u.People)

	fmt.Printf("\nTotal Tip: $%.2f\n", tip)
	fmt.Printf("Total Cost post tip and tax: $%.2f\n", total)
	fmt.Printf("Total cost per person: $%.2f\n\n", costPerPerson) // %.2f

}
