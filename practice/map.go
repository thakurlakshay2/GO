package main

import "fmt"



func main(){
	menu :=map[string]float64{
		"soup":4.99,
		"pie":7.2,
		"salad":3.3,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v:=range menu{
		fmt.Println(k," - ",v)
	}
}