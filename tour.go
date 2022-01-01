package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func createArray() {
	var a [2]string
	a[0] = "Hello"
	fmt.Println(a)
}

func createSlice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	a := names[0:2]
	b := names[1:3]
	a[0] = "XXX" // スライス内の要素を変更すると、元の配列内の要素も変更される
	b[1] = "YYY"
	fmt.Println(names)
}

func sliceLiteral() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func lengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[1:2]
	printSlice(s)
}

func arrayAppend() {
	var s []int

	s = append(s, 0)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func useRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

type MapStruct struct {
	Lat, Long float64
}

var m map[string]MapStruct

func useMap() {
	// pythonでいう辞書の作成
	m = make(map[string]MapStruct)
	m["Bell Labs"] = MapStruct{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

var n = map[string]MapStruct{
	"Bell Labs": MapStruct{
		40.68433, -74.39967,
	},
	"Google": MapStruct{
		37.42202, -122.08408,
	},
}

func mapOpe() {
	// map操作
	// nへ要素の挿入や更新
	n["Apple"] = MapStruct{
		14.533, -52.532,
	}
	// 要素の取得
	fmt.Println(n["Apple"])
	// 要素の削除
	delete(n, "Apple")
	// キーに要素が存在するかどうか
	elem, exist := n["Apple"]
	// もしnにAppleがあればexist変数はTrueになり、なければFalseになる
	// なおAppleが存在しない場合、elemはmapの要素の型のゼロ値となる
	fmt.Println(elem, exist)
}

// mapに渡すトップレベルの型が単純な型名である場合、リテラルの要素から推定できるのでその型名を省略することができる
// var n = map[string]MapStruct{
// 	"Bell Labs": {40.68433, -74.39967},
// 	"Google": {37.42202, -122.08408},
// }

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println(v1, v2, v3, p)
	createArray()
	createSlice()
	sliceLiteral()
	lengthAndCapacity()
	arrayAppend()
	useRange()
	useMap()
	fmt.Println(n)
	mapOpe()
}
