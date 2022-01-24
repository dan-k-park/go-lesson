package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{"nachos": 9.55, "calamari": 5.55},
		tip: 0,
	}

	return b
} 

// receiver function
// bill type is received denoted by variable b
// limits this function to bill objects
func (b *bill)format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		// the -25 forces the output to be 25 characters long
		// this is being used to line up the prices
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}


	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}

func (b *bill) updateTip(tip float64) {
	// if we receive a pointer, go automatically dereferences it
	b.tip = tip
} 

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}