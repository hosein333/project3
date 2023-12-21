package lib

import "errors"

var cityCode map[string]string = map[string]string{"011": "tehran jonob"}

func getCityName(code string)(string,error){
	val, ok := cityCode[code]
	if ok {
		return val, nil
	}
	return "", errors.New("City Dont Exist")
}
