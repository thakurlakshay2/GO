package main

import (
	"fmt"
	"strings"
)

//main function , first execution
func main(){
	fmt.Printf("hello")
	
	//String 
		// var name1 string = "string";
	name :="string"
		// fmt.Printf(name1)

		// fmt.Printf(name)

	//bits & memory (not really used a lot)
		// var num1 int8 = 4

		// fmt.Printf(num1);

		age :=18;
	fmt.Println("my age is",age,"my name is",name);
		//or just do this 
		// fmt.Println("my age is %v my name is %v",age,name);
		// fmt.Println("my age is %v my name is %v",age,name);

	//Arrays and
	var ages[3] int= [3]int{20,25,20};

	// var ages2=[3]int{1,2,3};
	fmt.Println(ages,len(ages));

		//no fixed length array
	var ages3=[]int{1,2,3};
	ages3[2]=84;
	fmt.Println(ages3)

	//slice
	ranger1 :=ages3[1:2];
	ranger2 := ages3[1:];
	ranger3 := ages3[:2];
	fmt.Println((ranger1))
	fmt.Println((ranger2))
	fmt.Println((ranger3))

	greetings :="hello with friends"
	fmt.Println(strings.Contains(greetings,"hello"))

}