package common

import (
	"encoding/json"
	"log"
)

func TestError(err error, errorText string) {
	if err != nil {
		log.Fatalf(errorText, err)
	}
}

func String2json(str string) map[string]interface{} {
	resMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(str), &resMap); err == nil {
		log.Println("string 2 json success" )
		return resMap
	} else {
		log.Fatalf("json string 2 map error, error: %v, string is: %s", err, str)
		return nil
	}

}


func Map2jsonString(m map[string]interface{}) string {
	mJson, _ := json.Marshal(m)
	mString := string(mJson)
	return mString
}