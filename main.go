package main 

import (
	"fmt"
	"project3/lib"
)

func main() {

	nID, err := lib.NewNationalID("0114022185")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("IsValid: ", nID.IsValid)
	city, err := nID.GetCity()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(city)

}
