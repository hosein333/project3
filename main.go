package main 

import "fmt"

type NationalID string

func (id NationalID) IsNationalIDlid(nID NationalID) bool {
	if len(nID) == 10 {
		for _, val := range nID {
		    _, err := strconv.Atoi(string(val))
		    if err != nil {
			return false
	}
}

		return true
	}
		return false
}


func main() {
	var nID NationalID
	nID = "127-221840"

	fmt.Println("My National ID Number %s is Vlid :%v", nID, nID.IsNationalIDlid(nID))
}