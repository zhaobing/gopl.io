package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
)

type Tdata struct {
	Name string
	Tags []string `json:"tags,omitempty"`
}

func main() {

	t := new(Tdata)
	t.Name = "test"
	//t.Tags = []string{"a","b"}

	toString, err := jsoniter.MarshalToString(t)
	if err != nil {
		logs.Error(err)
	}
	fmt.Println("", toString)
}
