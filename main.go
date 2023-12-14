package main 

import "fmt"

type NationalID string

func IsNationalIDlid(nID NationalID) bool {
	if len(nID) == 10 {
		return true
	}
		return false
}


func main() {
	var nID NationalID
	nID = "1273922184"

	fmt.Println("My National ID Number %s is Vlid :%v", nID, IsNationalIDlid(nID))
}
