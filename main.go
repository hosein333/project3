package main 

import (
	"fmt"
	"project3/lib"
)

func main() {

	nID1 := lib.NewNationalID("1274022185")

	fmt.Println("OutClass, IsValid: ", nID1.IsValid)

}
