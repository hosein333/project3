package citycode

import "errors"

var cityCode map[string]string = map[string]string{
	"011": "tehran jonob",
	"127": "esf",
}


func GetCityName(code string) (string, error){
	val, ok := cityCode[code]
	if ok {
		return val, nil
	}
	return "", errors.New("City Dont Exist")
}
