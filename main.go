package main 

import (
	"fmt"
	"project3/lib.go"
)

func main() {
	var nID1 ncode.NationalID
	nID1 = "127-221840"

	fmt.Printf("My National ID Number %s is Vlid :%v\n", nID1, nID1.IsVlid())

        var nID2 ncode.NationalID
        nID2 = "1274022185"

        fmt.Printf("My National ID Number %s is Vlid :%v\n", nID2, nID2.IsVlid())

}
