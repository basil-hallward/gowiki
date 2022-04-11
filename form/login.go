package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"regexp"
	"time"
	"crypto/md5"
	"io"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // urlが渡すオプションを解析する。POSTに対してはレスポンスパケットのボディを解析する。
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("Value: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // ここでwに書き込まれたものがクライアントに出力される
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // リクエストを取得
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("token -> ", token)
		t, _ := template.ParseFiles("./login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm() // フォームの内容を解析
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])

		token := r.Form.Get("token")
		if token != "" {
			fmt.Println(token)
		} else {
			fmt.Println("Token does not exist")
		}

		if email, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !email {
			fmt.Println("Email error")
		} else {
			fmt.Println("email: ", email)
		}

		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			// 数の変換でエラーが発生。つまり数字ではない。
			fmt.Println("getint error", err)
		} else {
			fmt.Println("getint: ", getint)
		}
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}