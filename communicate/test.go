package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func InputHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("input.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func OutputHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		t, err := template.ParseFiles("output.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, map[string]interface{}{"username": r.FormValue("username")})
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	//	http.HandleFunc("/input", InputHandler)
	//	http.HandleFunc("/output", OutputHandler)
	//	fmt.Println("服务端口:8000")                 //控制台输出信息
	//	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	//	if err != nil {
	//		log.Fatal("ListenAndServe: ", err)
	//	}
}
