package common

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

// Map data to structures based on sql tags in structures and convert types.
func DataToStructByTagSql(data map[string]string, obj interface{}) {
	objValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < objValue.NumField(); i++ {
		//Get the value corresponding to sql
		value := data[objValue.Type().Field(i).Tag.Get("sql")]
		//Get the name of the corresponding field
		name := objValue.Type().Field(i).Name
		//Get the corresponding field type
		structFieldType := objValue.Field(i).Type()
		//Get the type of the variable, or just write "string".
		val := reflect.ValueOf(value)
		var err error
		if structFieldType != val.Type() {
			//type conversion
			val, err = TypeConversion(value, structFieldType.Name()) //type conversion
			if err != nil {

			}
		}
		//Setting the type value
		objValue.FieldByName(name).Set(val)
	}
}

func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	//else if .......

	return reflect.ValueOf(value), errors.New("Unknown type:" + ntype)
}
