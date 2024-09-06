package main

import "fmt"

func updateName(x string){
	x="changed"
}

func updaeMenu(y map[string]float64){
	y["coffee"]=12.44
}
func types(){

	// GROUP A types : strings, ints , bools , floats , arrays , structs

	string1 :="initial"
	
	fmt.Println(string1) 
	updateName(string1);
	fmt.Println(string1) ; //No change even after function call

	//, when we pass a  variable directly go just created a copy for them , (for GROUP A only)
	//So all the "change" is done to the copy and when execution is done those copies are removed

	//GROUP B: slice , maps , functions (pass by reference)

	mappingObject := map[string]float64{
		"coffee":10.11,
		"water":0.00,
	}

	fmt.Println(mappingObject);
	updaeMenu(mappingObject);
	fmt.Println(mappingObject);

}