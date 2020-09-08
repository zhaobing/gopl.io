package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type SdkCfgCache struct {
	table map[string]*SdkCfg
}

type SdkCfg struct {
	ContainerCode string
	Country       string
}

func NewSdkCfg(containerCode string, country string) *SdkCfg {
	return &SdkCfg{ContainerCode: containerCode, Country: country}
}

func main() {
	//t1()
	t2()
}

func build(items interface{}) {

	vs := reflect.ValueOf(items)
	fmt.Println("", vs)
	kind := vs.Kind()
	fmt.Println("", kind)
	if kind != reflect.Slice && kind != reflect.Array {
		return
	}

	for i := 0; i < vs.Len(); i++ {
		v := vs.Index(i)
		if v.Kind() != reflect.Ptr {
			continue
		}
		fmt.Printf("v\t%s\t%T\t%s\t%s\n", v, v, v.Kind(), v.Elem())
		elem := v.Elem()
		elemKind := elem.Kind()
		t := elem.Type()
		fmt.Println("elemKind", elemKind, "t", t)

	}

	//for _, item := range items.([]interface{}) {
	//	fmt.Println("", item)
	//}
}

func t2() {
	cfgs := make([]*SdkCfg, 0)
	cfg1 := NewSdkCfg("极乐净土", "JP")
	cfg2 := NewSdkCfg("Numb", "US")
	cfgs = append(cfgs, cfg1, cfg2)
	fmt.Println("cfgs", cfgs)
	build(cfgs)
	//获取SdkCfg的类型名称
	//reflect.Indirect(cfg).Type().Name()
	//获取
}

func t1() {
	i := 3
	ofi := reflect.TypeOf(i)
	fmt.Printf(" %T\t%s\n", ofi, ofi.String())
	cfg := NewSdkCfg("极乐净土", "JP")
	ofCfg := reflect.TypeOf(cfg)
	of := reflect.ValueOf(cfg)
	fmt.Println("of", of)
	fmt.Printf(" %T\t%s\n", ofCfg, ofCfg.String())
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"
	fmt.Println("----------")
	v := reflect.ValueOf(3)      // a reflect.Value
	fmt.Println(v)               // "3"
	fmt.Printf("%T\t%v\n", v, v) // "3"
	fmt.Println(v.String())      // NOTE: "<int Value>"
	t := v.Type()                // a reflect.Type
	fmt.Println(t.String())      // "int"
}
