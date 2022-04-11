package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	// close関数によってchannelを閉じることができる
	// 閉じた後はいかなるデータも送信することはできない
	// v, ok := <-chという式でchannelがすでに閉じられているかテストすることができる
	// もしokにfalseが返ってきたら、channelはすでにどのようなデータもなく、閉じられているということになる
	// これは生産者側で閉じていることに注意する
	// また、channelはファイルのようなものではないので、頻繁に閉じる必要はない。何のデータも送ることがない場合、またはrangeループを終了させたい場合などでよい
	close(c)
}

func fibonacci_2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	go say("world")
	say("hello")
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:3], c)
	go sum(a[3:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	// channelは一時的に要素を保持することができる（バッファリング）
	// Goではchannelのサイズを指定することができる
	// make(chan type, value) でサイズを指定する

	// この例であればbチャンネルは、3個のint型飲み受け付け、そのほかはブロックする(エラー発生)
	// チャンネルには channel <- value で値を送信、hensuu <-channelでチャンネルから値を受け取る
	// チャンネルから値を受け取った時にその値はチャンネルから削除され、スペースが空く
	b := make(chan int, 3)
	b <- 1 // bチャンネルに1を送信
	b <- 5 // bチャンネルに5を送信
	b <- 10
	fmt.Println("b = ", <-b) // bチャンネルから値を受信
	fmt.Println("b = ", <-b)
	fmt.Println("b = ", <-b)

	// rangeによって、sliceやmapを操作するのと同じ感覚でバッファリング型のchannelを操作することができる
	d := make(chan int, 10)
	go fibonacci(cap(d), d)
	for i := range d {
		fmt.Println(i)
	}

	p := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-p)
		}
		quit <- 0
	}()
	fibonacci_2(p, quit)
}