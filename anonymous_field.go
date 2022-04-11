package main

import (
	. "fmt"
)

type Human struct {
	name string
	age int
	weight int
}

// 匿名フィールドはPythonでいう「継承」のようなもの
type Student struct {
	Human // Humanのフィールドを継承する
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	Println(mark.name)

	mark.age = 46
	Println(mark.age)
}