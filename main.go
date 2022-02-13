package main

import (
	"IM-gossip-space/ctrl"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

func RegisterView() {
	// note the directory here
	tpl, err := template.ParseGlob("view/**/*")
	if nil != err {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}

func main() {
	//bind the request and handler function
	//绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	//http.HandleFunc("/user/find", ctrl.FindUserById)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/contact/createcommunity", ctrl.CreateCommunity)
	//http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/chat", ctrl.Chat)
	//http.HandleFunc("/attach/upload", ctrl.Upload)
	//1 提供静态资源目录支持
	//http.Handle("/", http.FileServer(http.Dir(".")))

	//2 指定目录的静态文件
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))

	RegisterView()
	//RegisterIndex()

	fmt.Println("run at :8080")
	http.ListenAndServe(":8080", nil)

	//provide static resource support: File server
	//provide designated directory for encapsulation
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	RegisterView()

	http.ListenAndServe(":8080", nil)
}
