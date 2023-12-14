package main 

import "fmt"

type NationalID string

func main() {
	var nID NationalID
	nID = "1273922184"

	fmt.Println("My National ID Number :", nID)
}
