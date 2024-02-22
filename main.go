package main

import (
	"fmt"
	"project3/national_id"
)

func main() {

	nID, err := national_id.NewNationalID("1274022185")

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
