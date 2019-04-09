package util

import (
	"reflect"
	"strconv"
)

func InterfaceToInt(i interface{}) (int, error) {
	var (
		userID int
		err    error
	)
	userIDType := reflect.TypeOf(i)

	switch userIDType.String() {
	case "int":
		userID = i.(int)
	case "int8":
		userID = int(i.(int8))
	case "int16":
		userID = int(i.(int16))
	case "int32":
		userID = int(i.(int32))
	case "int64":
		userID = int(i.(int64))
	case "float32":
		userID = int(i.(float32))
	case "float64":
		userID = int(i.(float64))
	case "string":
		userIDStr := i.(string)
		userID, err = strconv.Atoi(userIDStr)
	}

	return userID, err
}
