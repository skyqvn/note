package main

import (
	"fmt"
	"reflect"
)

type stru struct {
	A int
	B string
}

func (st stru) C() {
	fmt.Println("this is func C")
}

func main() {
	var s = "hello go"
	var i = 100
	sptr := &s
	
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	t = v.Type()
	fmt.Println("t:", t, "v:", v)
	name := t.Name()
	kind := v.Kind()
	fmt.Println("name:", name, "kind:", kind, t.Kind())
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.Int16:
		fmt.Println("int16")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("chan")
	}
	fmt.Println("i is", v.Int())
	
	sv := reflect.ValueOf(sptr).Elem()
	fmt.Println(sv)
	fmt.Println("can set:", sv.CanSet())
	if sv.CanSet() {
		sv.Set(reflect.ValueOf("hello world"))
	}
	fmt.Println(sv)
	
	fmt.Println("can addr:", sv.CanAddr())
	if sv.CanAddr() {
		fmt.Println(sv.Addr())
	}
	
	sayFunc := func(say string) {
		fmt.Println(say)
	}
	sf := reflect.ValueOf(sayFunc)
	sf.Call([]reflect.Value{reflect.ValueOf("hello,this is sayFunc")})
	
	// 如果函数接收器是指针，记得TypeOf也用指针
	struv := stru{
		A: 19,
		B: "hello",
	}
	v = reflect.ValueOf(struv)
	t = reflect.TypeOf(struv)
	fmt.Printf("there are %d func in type stru\n", v.NumMethod())
	v.Method(0).Call(nil)
	f := v.MethodByName("C")
	f.Call(nil)
	fmt.Println(t.Method(0).Name)
	
	fmt.Println(v.NumMethod())
	fmt.Println(t.Field(0).Name)
	fmt.Println(v.Field(0))
	fmt.Println(v.FieldByName("B"))
}
