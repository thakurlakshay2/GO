package main

import (
	"fmt"
	"math"
	"sort"
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
		// fmt.Printf("my age is %v my name is %v",age,name);
		// fmt.Printf("my age is %v my name is %v",age,name);

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

	//SPLIT
	greetings :="hello with friends"
	fmt.Println(strings.Contains(greetings,"hello"))
	fmt.Println(strings.ReplaceAll(greetings,"hello","hi"))
	fmt.Println(strings.ReplaceAll(greetings,"hello","hi"))
	fmt.Println(strings.Split(greetings," "))

	//SORT
	sortArray :=[]int{45,20,35,30,75}
	sort.Ints(sortArray);
	fmt.Println(sortArray)
	fmt.Println(sort.SearchInts(sortArray,45))


	names :=[]string{"yoshi","lakshay","changes","mar"};
	fmt.Println((sort.SearchStrings(names,"nons")))

	//LOOPS
	loop :=0;
	for loop<5{
		fmt.Println("changes", loop)
		loop++
	}

	for i:=1; i<5;i++{
		fmt.Println(i)
	}

	for index , value :=range names{
		fmt.Printf("index %v value %v ",index,value);
	}
	fmt.Println()
	for _ , value :=range names{
		fmt.Printf("value %v ",value);
	}


	//Boolean & conditions
	data :=55

	if data<24 {
		fmt.Println("not found")
	}else {
		fmt.Println("age isse")
	}
	sayGreeting("nuuua")
	cycle(names,sayGreeting)
	fmt.Println(circleArea(4.3))
	fmt.Println(circleArea(4.6))
	fmt.Println(multipleReturns());



	Hello("lakshay")

	//MAP
	mapFunc()

	//TYPES
	types()

	//POINTERS
	pointers()

}

//FUNCTION
func sayGreeting(n string){
	fmt.Printf("number %v",n);
}
func cycle(n []string,fun func(string)){
	for _,v :=range n{
		fun(v);
	}
}
func circleArea(radius float64) float64{
	return math.Pi*radius*radius
}
//FUNCTION multiple returns
func multipleReturns()(string, string){
	return "lakshay","thakur"
}