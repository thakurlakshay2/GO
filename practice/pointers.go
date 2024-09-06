package main

import "fmt"

func updateUsingAddress(addressInfo *string){
	*addressInfo= "changed"
}
func referencePassing( referenceNo *string  ){
	fmt.Println(*referenceNo)
}
func pointers(){

	//pointers

	//basic '&' gives address where something is stored
	testing :="testing"

	fmt.Println("Address of testing varaible is : ",&testing)

	//Now if you are sending addrerss of something to a function or somewhere else, how does that place know that it is  a address/reference , 
	// or to read the value in that address instead of reading the address.

	//we use "*"
	m := &testing
	referencePassing(m)

	updateUsingAddress(m); // we send reference to the variable  


	fmt.Println(testing)  //now the change is made in the memory itself


}