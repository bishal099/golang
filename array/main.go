package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var dataList [8]string

	dataList[0] = "dfgh"
	dataList[7] = "dfghjk"

	fmt.Println(dataList)

	var dataList2 = [8]string{"fg", "ghj"}
	fmt.Print(dataList2)

}
