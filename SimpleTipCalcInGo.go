// Take args prices, tip. numOfPeople
// Return total per person

package main

import "fmt"

func TotalCost(Price, Tax, Tip float64, People uint) float64 {
	var taxConverted float64 = Tax / 100.0
	var total float64 = Price + (1 + taxConverted)
	var CostPerPerson float64 = total / float64(People)
	return CostPerPerson
}

func main() {
	//args
	/* 	var Price = flag.Float64("price", 0.0, "Price of goods")
	   	var Tax = flag.Float64("Tax", 0.0, "State food tax")
	   	var Tip = flag.Float64("tip", 0.0, "Percentage to Tip")
	   	var People = flag.Uint("people", 0, "Number of people") */
	var Price float64 = 10.0
	var Tax float64 = 8.5
	var Tip float64 = 15.0
	var People uint = 3

	var CostPerPerson = TotalCost(Price, Tax, Tip, People)

	fmt.Printf("Total cost per person is %v.2f \n", CostPerPerson)

}
