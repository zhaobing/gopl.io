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
	ContainerCode string `common_cache:"key"`
	Country       string `common_cache:"key"`
	VerCode       int32  `common_cache:"key"`
	LimitNum      int
}

func NewSdkCfg(containerCode string, country string, verCode int32, limitNum int) *SdkCfg {
	return &SdkCfg{ContainerCode: containerCode, Country: country, VerCode: verCode, LimitNum: limitNum}
}

func main() {
	//t1()
	t2()
	//t3()
}

func t3() {
	Register(new(SdkCfg), new(SdkCfgCache))
}

//注册类型
func Register(models ...interface{}) {
	for _, model := range models {
		registerModel(model)
	}
}

func registerModel(model interface{}) {
	val := reflect.ValueOf(model)
	typ := reflect.Indirect(val).Type()

	fmt.Println("", model, val, val.Kind(), typ, typ.Kind())

	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("<common_cache.RegisterModel> cannot use non-ptr model struct `%s`", getFullName(typ)))
	}

	// registerModel(&u)
	if typ.Kind() == reflect.Ptr {
		panic(fmt.Errorf("<common_cache.RegisterModel> only allow ptr model struct, it looks you use two reference to the struct `%s`", typ))
	}
}

// get reflect.Type name with package path.
func getFullName(typ reflect.Type) string {
	return typ.PkgPath() + "." + typ.Name()
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
		fmt.Println("elemKind", elem.Kind(), "elemType", elem.Type())
		//获取字段值和字段的注解
		// Build map of fields keyed by effective name.
		for i := 0; i < elem.NumField(); i++ {
			fieldInfo := elem.Type().Field(i) // a reflect.StructField
			tag := fieldInfo.Tag
			name := fieldInfo.Name
			fieldValue := elem.Field(i)
			fmt.Println("fieldName", name, "tag", tag, "fieldValue", fieldValue)
		}
	}

	//for _, item := range items.([]interface{}) {
	//	fmt.Println("", item)
	//}
}

func t2() {
	cfgs := make([]*SdkCfg, 0)
	cfg1 := NewSdkCfg("极乐净土", "JP", 1, 666)
	cfg2 := NewSdkCfg("Numb", "US", 2, 333)
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
	cfg := NewSdkCfg("极乐净土", "JP", 1, 333)
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
