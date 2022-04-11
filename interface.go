package main

import (
	"fmt"
	"strconv"
)

// 複数のクラスで共通のメソッドがある場合、インターフェイスでまとめて呼び出すことができる

// 人間とカメで、それぞれ「食べる」という行動に必要なメソッドを用意して、インターフェースでまとめて呼び出す

// 食べるためのインターフェース
type Eater interface {
	PutIn() // 口に入れる
	Chew() // 噛む
	Sallow() // 飲み込む
}

// 人間の構造体
type Human struct {
	Height int // 身長
}

// 亀の構造体
type Turtle struct {
	King string // 種類
}

// 人間用のメソッド
func (h Human) PutIn() {
	fmt.Println("道具を使って口に運ぶ")
}
func (h Human) Chew() {
	fmt.Println("歯でしっかりかむ")
}
func (h Human) Sallow() {
	fmt.Println("よく噛んだら飲み込む")
}

// 亀用のインターフェース
func (t Turtle) PutIn() {
	fmt.Println("獲物を見つけたら首をすばやく伸ばして噛む")
}
func (t Turtle) Chew() {
	fmt.Println("クチバシで噛み砕く")
}
func (t Turtle) Sallow() {
	fmt.Println("小さく砕いたら飲み込む")
}

// インターフェースが引数となる、「食べる」メソッド
func EatAll(e Eater) {
	e.PutIn()
	e.Chew()
	e.Sallow()
}

func main() {
	// インターフェースを通じて複数のメソッドを一度に呼び出すことができる
	// クラスが複数あっても、共通するメソッドがあればインターフェースで一つにまとめることができる
	man := Human{Height: 170}
	cheloniaMydas := Turtle{King: "アオウミガメ"}

	fmt.Println("<人間が食べる>")
	EatAll(man)

	fmt.Println("<カメが食べる>")
	EatAll(cheloniaMydas)

	// 空のinterfaceは任意の型の値を保存することができる
	var empty interface{}
	var i int = 5
	s := "Hello World"
	empty = i
	empty = s
	fmt.Println(empty)

	list :=  make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		switch value := element.(type) {
			case int:
				fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
			case string:
				fmt.Printf("list[%d] is a string and value is %s\n", index, value)
			case Person:
				fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
			default:
				fmt.Printf("list[%d] is of a different type", index)
		}
	}
}

// interface変数に保存されているオブジェクトの型を知るには
type Element interface{}
type List [] Element

type Person struct {
	name string
	age int
}

// Stringメソッドを定義（fmt.Stringer）
func (p Person) String() string {
	return "{name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}