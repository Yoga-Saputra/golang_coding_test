package helper

import (
	"encoding/json"
	"strconv"
)

func InterfaceToString(params interface{}) string {
	byte, _ := json.Marshal(params)
	return string(byte)
}

func IsPalindrome(str string) string {
	result := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	var data string

	if str == string(result) {
		data = "Palindrome"
	} else {
		data = "Not Palindrome"
	}
	return data
}

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert interface to a json string

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map

	if err != nil {
		return newMap, err
	}

	return
}

func GetApeared(params interface{}) int {
	byteParam, _ := json.Marshal(params)
	appeared, _ := strconv.Atoi(string(byteParam))

	return appeared
}

func GetSliceCreated(params interface{}) []string {
	byteCreated, _ := json.Marshal(params) // Convert interface to a json string
	var sliceCreated []string
	err := json.Unmarshal(byteCreated, &sliceCreated) // Convert to []string
	if err != nil {
		panic(err)
	}

	return sliceCreated
}

func GetInfluenced(params interface{}, value string) []string {
	var sliceInfluencedBy []string
	influencedInterface := params.(map[string]interface{})[value].([]interface{})
	byteInfluencedBy, _ := json.Marshal(influencedInterface)    // Convert interface to a json string
	err := json.Unmarshal(byteInfluencedBy, &sliceInfluencedBy) // Convert to []string
	if err != nil {
		panic(err)
	}

	return sliceInfluencedBy
}
