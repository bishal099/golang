package main

import "fmt"

//declare global variable
const GlobalVariable = "global value here"

func main() {

	//var type ckeck
	var username string = "bishal"
	fmt.Println(username)
	fmt.Printf("Variable Type : %T \n", username)

	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("Variable Type : %T \n", isTrue)

	//access global variable
	print(GlobalVariable)

	//implicit var
	var testVar = "test var declaration"
	fmt.Println(testVar)

	//novar
	totalPopulation := 45678934567890
	fmt.Println(totalPopulation)

}
