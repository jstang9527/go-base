package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//首页
func hellowWorld(reponse http.ResponseWriter, req *http.Request) {
	// str := `<h1 style="color:red;text-align:center">Golang</h1>`
	b, err := ioutil.ReadFile("./template/index.html")
	if err != nil {
		fmt.Println("读取html失败")
		reponse.Write([]byte("服务器出错..."))
		return
	}
	fmt.Printf("[Info] %v %v get[%v]\n", time.Now().Format("2006-01-02 15:04:05"), req.Host, req.URL)
	reponse.Write(b)

}

//图片获取
func getImages(reponse http.ResponseWriter, req *http.Request) {
	url := strings.Split(req.URL.RequestURI(), "/")
	imgName := fmt.Sprintf("./images/%v", url[len(url)-1])
	b, err := ioutil.ReadFile(imgName)
	if err != nil {
		fmt.Println("读取图片失败")
		reponse.Write([]byte("服务器出错..."))
		return
	}
	fmt.Printf("[Info] %v %v get[%v]\n", time.Now().Format("2006-01-02 15:04:05"), req.Host, req.URL)
	reponse.Write(b)
}

//方法
func temp(reponse http.ResponseWriter, req *http.Request) {
	fmt.Printf("[Info] %v %v get[%v]\n", time.Now().Format("2006-01-02 15:04:05"), req.Host, req.URL)
	// fmt.Println(ioutil.ReadAll(req.Body))

	url := strings.Split(req.URL.RequestURI(), "/")
	endDuan := url[len(url)-1]
	if req.Method == "GET" {
		//1.先判断是否有?号
		if strings.Contains(endDuan, "?") {
			xs := strings.Split(endDuan[1:], "&")
			var strSlice string
			for i, v := range xs {
				array := strings.Split(v, "=")
				strSlice += fmt.Sprintf("%v=<%d>=%v\n", array[0], i, array[1])
			}
			reponse.Write([]byte(strSlice))
			return
		}
		reponse.Write([]byte("GET"))
		return
	}
	if req.Method == "POST" {
		reponse.Write([]byte("POST"))
		return
	}

	reponse.Write([]byte("Other"))
}

func main() {

	http.HandleFunc("/hello", hellowWorld)
	http.HandleFunc("/images/", getImages)
	http.HandleFunc("/xxx/", temp)
	if err := http.ListenAndServe("0.0.0.0:10001", nil); err != nil {
		fmt.Println("创建HTTP服务器失败, err: ", err)
		return
	}
	fmt.Println("创建HTTP服务器成功败")
}
