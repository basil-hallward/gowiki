package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Goにはクラスはないが、メソッドを定義することができる
// メソッドはレシーバという特別な引数を取る
func (v Vertex) Abs() float64 {
	// レシーバはfuncキーワードとメソッド名の間に自身の引数をリストで表現する
	// この例ではAbsメソッドはvという名前のVertex型のレシーバを持つことを意味する
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// こちらは通常の関数(機能は同じ)
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scala(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// メソッド呼び出し
	fmt.Println(v.Abs())

	// 通常の関数呼び出し
	fmt.Println(Abs2(v))

	v.Scala(10)
	fmt.Println(v.Abs())
}
