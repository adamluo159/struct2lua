package struct2lua

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func spaceLayer(layer int) string {
	var result string = ""
	for i := 0; i < layer; i++ {
		result += "\t"
	}
	return result
}

func ToLuaObject(layer int, i interface{}) string {
	var result string = ""

	fmt.Println(i)
	k := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	var kName string = ""
	switch v.Kind() {
	case reflect.Struct:
		var split = ""
		if v.NumField() > 1 {
			split = ","
		}
		if layer > 0 {
			result += "{"
		}
		if layer == 0 {
			split = ""
		}

		prefix := "\n" + spaceLayer(layer)

		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).CanInterface() == false {
				break
			}
			keyName := k.Field(i).Name
			subV := v.Field(i).Interface()
			result += prefix + keyName + " = " + ToLuaObject(layer+1, subV) + split

		}
		if layer > 0 {
			result += "\n" + spaceLayer(layer-1) + "}"
		}

	case reflect.Slice:
		result = "{\n"
		for i := 0; i < v.Len(); i++ {
			result += spaceLayer(layer) + ToLuaObject(layer+1, v.Index(i).Interface()) + ",\n"
		}
		result += spaceLayer(layer-1) + "}"
	case reflect.String:
		s := v.String()
		if s == "" {
			result = kName + "nil"
		} else {
			result = kName + "\"" + s + "\""
		}

	case reflect.Int:
		vInt := (int)(v.Int())
		str := strconv.Itoa(vInt)
		result = kName + str

	case reflect.Bool:
		var b bool = v.Bool()
		result = kName + strconv.FormatBool(b)
	case reflect.Map:
		result = "{\n"
		for _, vmap := range v.MapKeys() {
			i := v.MapIndex(vmap).Interface()
			result += spaceLayer(layer) + vmap.String() + " = "
			result += ToLuaObject(layer+1, i) + ",\n"

		}
		result += spaceLayer(layer-1) + "}"

	default:
		result = "nil"
	}

	return result
}

func ToLuaConfig(fileName string, obj interface{}) bool {
	k := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		fmt.Println("this data is not struct!!")
		return false
	}

	layer := 0
	head := spaceLayer(layer) + k.Name() + "= {\n"

	for i := 0; i < k.NumField(); i++ {
		//k1 := k.Field(i)
		//head += ToLuaObject(layer+1, &k1, v.Field(i)) + ",\n"
	}
	head += spaceLayer(layer) + "}"

	f, err := os.Create(fileName + ".lua")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer f.Close()
	f.WriteString(head)

	return true
}
