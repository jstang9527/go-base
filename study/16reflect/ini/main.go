package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//MysqlConfig x
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

//RedisConfig x
type RedisConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

//Config x
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//0.data必须为指针
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("data param should be a pointer") //格式化输出之后返回一个error
		return
	}
	//1.data必须是struct
	if t.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("data param should be a struct") //格式化输出之后返回一个error
		return
	}
	//2.读取文件
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\n")
	var structName string
	//3.一行行读
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line) //去掉首位空格
		//3.1注释/空行则跳过
		if len(line) == 0 || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//3.2'[]'节的断定
		if strings.Contains(line, "[") || strings.Contains(line, "]") {
			//3.2.1 无效的节
			if line[len(line)-1] != ']' { //最后一个字符非 ‘]’ 结尾
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			if len(line) < 3 { //[]中间无内容
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			sectionName := line[1 : len(line)-1]
			// 寻找是哪个数据库结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if field.Tag.Get("ini") == sectionName {
					structName = field.Name
					fmt.Printf("外包结构体:%#v\n", structName)
					break
				}
			}
			continue
		}
		//3.3 key-value
		if len(line) < 3 { //起码得3个字节,d=3
			err = fmt.Errorf("line:%d syntax error", idx+1)
			return
		}
		keyValue := strings.Split(line, "=")
		if len(keyValue) != 2 { //非key-value
			err = fmt.Errorf("line:%d syntax error", idx+1)
			return
		}
		//3.4根据strcutName把具体的数据库结构体取出来
		v := reflect.ValueOf(data)
		sValue := v.Elem().FieldByName(structName) //获取结构体的值信息
		sType := sValue.Type()                     //获取结构体的类型信息
		if sType.Kind() != reflect.Struct {
			err = fmt.Errorf("data中的%s字段必须为一个结构体", structName)
			return
		}
		var fieldName string       //存储key
		var fieldType reflect.Type //值类型
		//3.5遍历结构体每一个字段，判定值类型
		for i := 0; i < sType.NumField(); i++ {
			field := sType.Field(i)
			if strings.ToLower(field.Tag.Get("ini")) == keyValue[0] {
				fieldName = field.Name
				fieldType = field.Type
				break
			}
		}
		if len(fieldName) == 0 { //结构体没有对应的值
			continue
		}
		// fmt.Printf("%v, %T,%#v\n", fieldName, fieldType)
		//4.对结构体进行赋值
		filedObj := sValue.FieldByName(fieldName)
		switch fieldType.Kind() {
		case reflect.String:
			filedObj.SetString(keyValue[1])
		case reflect.Int:
			var vInt int64
			vInt, err = strconv.ParseInt(keyValue[1], 10, 64)
			if err != nil {
				err = fmt.Errorf("line:%d type error, err:%v", idx+1, err)
				return
			}
			filedObj.SetInt(vInt)
		default:
			fmt.Println("none")
		}
	}
	return
}

func main() {
	var mc Config
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		fmt.Println("load ini failed, err:", err)
		return
	}
	fmt.Println(mc.MysqlConfig)
	fmt.Println(mc.RedisConfig)

}
