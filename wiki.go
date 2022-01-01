package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	// Pageへのポインタであるpをレシーバとして受け取るsaveという名前のメソッド
	// これはパラメータを取らず、error型を返す
	// エラー値を返すのはファイルの書き込み中に何か問題が発生した場合に、アプリケーション側で処理できるようにするため
	// 問題がなければPage.save()はnil(ゼロ値)を返す

	filename := p.Title + ".txt"

	// 3番目のパラメータとして渡される8進整数リテラル0600は、現在のユーザのみの読み書き可能なファイルを作成することを意味する
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ = loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
