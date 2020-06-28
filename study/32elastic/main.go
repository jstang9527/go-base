package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

// Person 表
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.231.139:9200"))
	if err != nil {
		fmt.Println("connect to es failed, err: ", err)
		return
	}
	fmt.Println("connect to es success")
	p1 := Person{Name: "rion", Age: 22, Married: false}
	put1, err := client.Index().Index("user").Type("person").BodyJson(p1).Do(context.Background())
	if err != nil {
		fmt.Println("put data to es failed, err: ", err)
	}
	fmt.Println("put data to es success.", put1.Id, put1.Index, put1.Type)
}

func getRecord() {

}
