package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// Widget x
type Widget interface {
	ID() string
}

// ID 接口Widget的方法被widget结构体实现
func (w widget) ID() string {
	return w.id
}

type widget struct {
	id string
}

// NewWidget x
func NewWidget() Widget { //返回的是一个接口，只能调用接口的函数
	return widget{
		id: uuid.NewV4().String(),
	}
}

func main() {
	fmt.Println("Hello, World!")
	obj := NewWidget()
	fmt.Printf("%v\n", obj.ID())
}
