package main

import (
	"fmt"
	"strings"
	"unicode"
)

func baseHandler() {
	s1 := "小明"
	s2 := "去上学"
	//拼接字符串
	value1 := s1 + s2
	fmt.Printf("直接相加拼接: %#v\n", value1)
	value2 := fmt.Sprintf("%s%s", s1, s2)
	fmt.Printf("通过fmt拼接回调输出: %s\n", value2)

	//字符串长度v
	fmt.Printf("用len函数求 %#v byte字节数量为:%v\n", value1, len(value1)) //15字节,一个汉字3字节

	//分割字符串
	value3 := strings.Split(value1, "去")
	fmt.Printf("%#v被'去'字符分割后:%s\n", value1, value3)

	//包含
	bool1 := strings.Contains(value1, "小明")
	fmt.Printf("%#v是否包含小明: %v\n", value1, bool1)

	//前后缀判定
	bool2 := strings.HasPrefix(value1, "小明")
	fmt.Printf("%#v是否以小明为前缀: %v\n", value1, bool2)
	bool3 := strings.HasSuffix(value1, "吃饭")
	fmt.Printf("%#v是否以吃饭为后缀: %v\n", value1, bool3)

	//判定字串的位置
	enStr := "abcdefgd"
	index := strings.Index(enStr, "d") //最先出现
	lastIndex := strings.LastIndex(enStr, "d")
	fmt.Printf("%#v字符串中'd'的首次出现下标位置为:%d, 最后出现的位置为:%d\n", enStr, index, lastIndex)

	//join操作
	var b2 = []string{"go", "lang", "!"}
	joinStr := strings.Join(b2, "##")
	fmt.Printf("join操作,往数组%#v值间插入字符串'##': %#v\n", b2, joinStr)
}

func curdHandler() {
	/*类型别名
	uint8类型, 或称byte型，表示ACSII码的一个字符
	int32类型, 或称rune类型, 代表一个UTF-8字符, 中文及其他为rune类型
	go语言中为了处理非ACSII码类型的字符，定义了新的rune类型
	*/

	//从字符串中拿取具体字符
	strEnCn := "你好, golang."
	fmt.Printf("字符串%#v的字节长度为:%d\n下标\t字符\t字符字节\n", strEnCn, len(strEnCn))

	for k, v := range strEnCn {
		fmt.Printf("%d\t%c\t%v\n", k, v, len(string(v))) //用字符占位符表示
	}

	//修改字符串
	//需要先转byte或rune类型, 完成后再转为string, 转换都需要重新分配内存, 并复制字节数组
	strEnCn2 := []rune(strEnCn)
	strEnCn2[0] = '红' //单引号包裹的字符
	strEnCn2 = append(strEnCn2, '棒')
	fmt.Printf("字符数组: %c, 类型:%T, 值类型:%T\n", strEnCn2, strEnCn2, strEnCn2[0])
	newStr := string(strEnCn2)
	fmt.Printf("修改后的字符串: %#v\n", newStr)
}

func homeWork() {
	//统计中英文字符串中汉字的数量
	s1 := "你好,goland!"
	s2 := []rune(s1)
	var num1 int8
	for _, v := range s2 {
		if len(string(v)) != 1 { //字符转string,再计算字节长度
			num1++
		}
	}
	fmt.Printf("字符串%#v中汉字有%d个\n", s2, num1)

	//标准，统计中英文字符串中汉字的数量
	var num int8
	for _, v := range s2 {
		if unicode.Is(unicode.Han, v) {
			fmt.Printf("%c", v)
			num++
		}
	}
	println(num)
}

func homeWork2() {
	//统计单词出现频率
	str := "how do you do"
	s1 := strings.Split(str, ` `)
	var myMap = make(map[string]int, 10)
	for _, v := range s1 {
		myMap[v]++
	}
	fmt.Println(myMap)
}

func homeWork3() {
	//回文判定
	//str := "上海自来水水来自海上" //10位
	str := "上海自来水来自海上" //9位
	str2 := []rune(str)
	chartLength := len(str2)
	var middleIndex int
	var flag bool
	if chartLength%2 != 0 {
		middleIndex = chartLength/2 + 1
	} else {
		middleIndex = chartLength / 2
	}

	fmt.Println(middleIndex)
	for i := chartLength - 1; i >= middleIndex; i-- {
		fmt.Println(str2[i])
		if str2[i] != str2[chartLength-i-1] {
			fmt.Println("No")
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("yes")
	}
	fmt.Println()
}

func main() {
	homeWork3()
	fmt.Println()
}
