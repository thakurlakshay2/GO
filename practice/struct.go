package main

import "fmt"


type Bill struct{
	name string
	items map[string]float64
	tip float64
}
func makeNewBill(name string) Bill{
	b := Bill{
		name:name,
		items: map[string]float64{
			"water":22.22,
			"pasta":200.22,
			"shake":140.00,
		},
		tip: 50.22,
	}
	return b
}

//basically anyone can call format bill function , so we need to associalte bill to the function

//Not like this
// func formatBill( billData Bill ){

// }
//We now receive bill object , by doing this we have now limited this function to only bill objects

//this example is of pass bas reference
func (b *Bill) formatBill() string{
	fs := "Bill Breakdown \n"
	total := 0.0

	for k, v :=range b.items{
		fs += fmt.Sprintf("%-25v...%v \n",k+":",v)
		total+=v
	}

	fs+=fmt.Sprintf("%-25v ...$%0.2f","total:",total)
	return fs
}

//This is pass by value , any change made to it , will not be reflected in the actual struct
func (b Bill) formatBillByValue() string{
	fs := "Bill Breakdown \n"
	total := 0.0

	for k, v :=range b.items{
		fs += fmt.Sprintf("%-25v...%v \n",k+":",v)
		total+=v
	}

	fs+=fmt.Sprintf("%-25v ...$%0.2f","total:",total)
	return fs
}

