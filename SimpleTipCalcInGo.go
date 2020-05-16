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
	var CostPerPerson float64 = totalTot / float64(People)
	return totalTip, totalTot, CostPerPerson
}

func main() {
	//args
	var Price = flag.Float64("Price", 100.0, "Price of goods")
	var Tax = flag.Float64("Tax", 8.5, "State food tax")
	var Tip = flag.Float64("Tip", 15.0, "Percentage to Tip")
	var People = flag.Uint("People", 3, "Number of people")
	flag.Parse()

	var tip, total, CostPerPerson = TotalCost(*Price, *Tax, *Tip, *People)

	fmt.Printf("\nTotal Tip: $%.2f\n", tip)
	fmt.Printf("Total Cost post tip and tax: $%.2f\n", total)
	fmt.Printf("Total cost per person: $%.2f\n\n", CostPerPerson) // %.2f

}
