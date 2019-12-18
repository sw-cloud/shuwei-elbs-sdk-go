package utils

import "encoding/json"

func ToJsonStr(v interface{}) string  {
	data, err := json.Marshal(v)
	if err != nil{
		return  ""
	}
	return string(data)
}
