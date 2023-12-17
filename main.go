package main 

import (
	"fmt"
	"project3/lib.go"
)

func main() {
	var nID1 ncode.NationalID
	nID1.StringFormat = "1273588291"

	fmt.Printf("My National ID Number %s is Vlid :%v\n", nID1.StringFormat, nID1.IsVlid())

        var nID2 ncode.NationalID
        nID2.StringFormat = "1274022184"

        fmt.Printf("My National ID Number %s is Vlid :%v\n", nID2.StringFormat, nID2.IsVlid())

}
