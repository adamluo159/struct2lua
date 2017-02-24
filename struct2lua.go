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
			split = "\n"
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

func ToLuaConfig(fileName string, Id int, obj interface{}) bool {
	idStr := strconv.Itoa(Id)

	head := fileName + " = "
	head += ToLuaObject(1, obj)
	f, err := os.Create(fileName + idStr + ".lua")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer f.Close()
	f.WriteString(head)

	return true
}
