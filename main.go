package main 

import (
	"fmt"
	"project3/lib.go"
)

func main() {
	var nID1 lib.NationalID
	nID1.StringFormat = "1274022184"

	fmt.Printf("My National ID Number %s is Vlid :%v\n", nID1.StringFormat, nID1.IsVlid())

	nID2 := &lib.NationalID{StringFormat : "1274022185"}

        fmt.Printf("My National ID Number %s is Vlid :%v\n", nID2.StringFormat, nID2.IsVlid())

	fmt.Println("OutClass, IsValid: ", nID2.IsValid)

}
