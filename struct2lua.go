package struct2lua

import (
	"fmt"
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

func toLuaObject(layer int, k *reflect.StructField, v reflect.Value) string {
	var result string = ""

	var kName string = spaceLayer(layer)
	if k != nil {
		kName = kName + k.Name + " = "
	}

	switch v.Kind() {
	case reflect.Struct:
		result += kName + "{"
		prefix := "\n"
		for i := 0; i < v.NumField(); i++ {
			subK := k.Type.Field(i)
			result += prefix + toLuaObject(layer+1, &subK, v.Field(i))
			prefix = ",\n"
		}
		result += "\n" + spaceLayer(layer) + "}"

	case reflect.Slice:
		result += kName + "{"
		prefix := "\n"

		for i := 0; i < v.Len(); i++ {
			result += prefix + toLuaObject(layer+1, nil, v.Index(i))
			prefix = ",\n"
		}
		result += "\n" + spaceLayer(layer) + "}"

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

	default:
		result = "nil"
	}

	return result
}

func ToLuaConfig(obj interface{}) string {
	k := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		fmt.Println("this data is not struct!!")
		return "false"
	}

	layer := 0
	head := spaceLayer(layer) + k.Name() + "= {\n"

	for i := 0; i < k.NumField(); i++ {
		k1 := k.Field(i)
		head += toLuaObject(layer+1, &k1, v.Field(i)) + ",\n"
	}
	head += spaceLayer(layer) + "}"
	return head
}
