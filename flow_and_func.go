package main

import "fmt"

// 可変長引数
func myFunc(arg ...int) {
	for _, n := range arg {
		fmt.Println("And the number is: %d\n", n)
	}
}

// 値渡しと参照渡し
// 引数に値を渡すとき、実際にはその値のコピーが渡される
func add1(a int) int {
	a++
	return a
}

// xそのものを渡したくなった場合は？
// ポインタを利用する。
// 我々は変数がメモリの中にある特定の位置に存在していることを知っている
// 変数を修正するということは、つまり変数のアドレスにあるメモリを修正しているということになる。
// xの存在するアドレスである&xを関数に渡し 引数の型をポインタ変数である*intに変更する
func add2(a *int) int {
	*a = *a+1
	return *a
}
/*
ポインタの特徴
・複数の関数が同じオブジェクトに対して操作を行うことができる
・ポインタ渡しは比較的軽い（8バイト）。ただのメモリアドレス。ポインタを使って大きな構造体を渡すことができる。
・channel、slice、mapの３つの型はメカニズムを実現するポインタのようなものなので、直接渡すことができる
*/

func main() {
	myFunc()

	x := 3
	x1 := add1(x) // x変数をadd1の引数に渡す
	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x) // x変数には何の変化も発生しない

	x2 := add2(&x)
	fmt.Println("x+1 = ", x2)
	fmt.Println("x = ", x) // x自体も変化している
}