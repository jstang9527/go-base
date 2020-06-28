package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int
}

func structToJSON(v interface{}) {
	var data = make(map[string]interface{}, 10)
	key := reflect.TypeOf(v)
	if key.Kind() != reflect.Struct {
		panic("传参非结构体.")
	}
	value := reflect.ValueOf(v)
	//取key字段
	for i := 0; i < key.NumField(); i++ {
		field := key.Field(i)
		if jsonKey := field.Tag.Get("json"); jsonKey != "" {
			if jsonValue := value.FieldByName(field.Name); true {
				fmt.Println(jsonValue)
				data[jsonKey] = jsonValue
			} else {
				data[jsonKey] = nil
				fmt.Println("dad")
			}

		} else {
			jsonKey = field.Name
			data[jsonKey] = value.FieldByName(jsonKey)
			if jsonValue := value.FieldByName(field.Name); true {
				fmt.Println(jsonValue)
				data[jsonKey] = jsonValue
			} else {
				data[jsonKey] = nil
				fmt.Println("dad")
			}
		}
	}
	//取value字段

	fmt.Println(data)

}

func main() {
	str1 := student{
		Name: "jack",
		// Score: 90,
	}
	// t := reflect.TypeOf(str1)
	// //通过for循环遍历结构体所有字段信息
	// fmt.Println(t.NumField())
	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)
	// 	// fmt.Println(field)
	// 	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	// 	// break
	// }
	structToJSON(str1)
}
