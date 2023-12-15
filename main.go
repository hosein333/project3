package main 

import (
	"fmt"
	"project3/ncode"
)

func main() {
	var nID ncode.NationalID
	nID = "127-221840"

	fmt.Printf("My National ID Number %s is Vlid :%v", nID, nID.IsVlid(nID))
}
