package main

import (
	"fmt"
	"reflect"
)

// interfaceにもStructの匿名フィールドのようなロジックを導入することができる
// もしinterface1がinterface2の組み込みフィールドであればinterface2は自動的にinterface1のメソッドを含むことになる
// type Interface interface {
// 	sort.Interface
// 	Push(x interface{})
// 	Pop() interface()
// }

func print(v interface{}) {
	tv := reflect.TypeOf(v)
	fmt.Println("Kind: ", tv.Kind(), "Name: ", tv.Name())
}

func main() {
	stringValue := "test"
	intValue := "test"

	print(stringValue)
	print(intValue)
	print(&stringValue)
}