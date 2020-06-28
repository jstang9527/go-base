package main

import (
	"fmt"
	"reflect"
)

func reflectType(v interface{}) {
	value := reflect.TypeOf(v)
	fmt.Printf("类型:%v, 类型名:%v, 类型种类:%v\n", value, value.Name(), value.Kind())
}

func reflectValue(v interface{}) {
	value := reflect.ValueOf(v)
	kind := value.Kind()
	switch kind {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("int64: %d\n", int64(value.Int()))
	case reflect.Float32:
		fmt.Printf("float32: %.2f\n", float32(value.Float()))
	case reflect.Float64:
		fmt.Printf("float64: %.2f\n", float64(value.Float()))
	}
}

func reflectSetValue(v interface{}) {
	addr := reflect.ValueOf(v)
	kind := addr.Elem().Kind()
	switch kind {
	case reflect.Int64:
		//value.SetInt(200)
		//反射中使用Elem()方法获取指针对应的值
		addr.Elem().SetInt(200)
	case reflect.Float64:
		// value.SetFloat(86.32)
		addr.Elem().SetFloat(400)
	}
}

//isNil常用于判断指针是否为空
//isValid常用于判断返回值是否有效
func nilOrValid() {
	var a *int // *int 类型空指针
	fmt.Println("var a *int isNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("nil isValid:", reflect.ValueOf(nil).IsValid())

	b := struct{}{} //匿名结构体
	//从匿名结构体中查找ab字段
	fmt.Println("不存在结构体成员:", reflect.ValueOf(b).FieldByName("ab").IsValid())
	//从匿名结构体中查找ab方法
	fmt.Println("不存在结构体方法:", reflect.ValueOf(b).MethodByName("ab").IsValid())

	//map
	c := map[string]int{}
	//从map中查找不存在的键
	fmt.Println("map中不存在的键", reflect.ValueOf(c).MapIndex(reflect.ValueOf("ab")).IsValid())
}

type person struct {
	name string
}

func main() {
	// p := person{name: "dd"}
	// reflectType(p)
	// reflectValue(float64(32))
	// var a float64 = 3.14
	// var b int64 = 999
	// reflectSetValue(&a)
	// reflectSetValue(&b)
	// fmt.Printf("%T-%T,%v-%v", a, b, a, b)
	nilOrValid()
}
