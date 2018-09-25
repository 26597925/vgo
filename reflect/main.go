package main

import(
"fmt"
"reflect"
)

func main() {
	d := &Test{"lyl", 18}
	v := reflect.ValueOf(d)
	ele := v.Elem()
	t := ele.Type()

	fmt.Println("读取这个对象的所有属性")
	for i:=0; i<t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, ele.Field(i).Interface())
	}

	fmt.Println("读取所有对象的方法")
	for i:=0; i <t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}

	fmt.Println("函数调用测试")
	ele.MethodByName("Stringtest").Call(nil)

	fmt.Println("未修改的年龄", d.Age)
	fmt.Println("待参函数调用")
	iage := reflect.ValueOf(11)
	v.MethodByName("Stringtest1").Call([]reflect.Value{iage})

	fmt.Println("反射调用通过指针修改后的年龄", d.Age)
	
}