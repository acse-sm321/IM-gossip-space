package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func user_login(writer http.ResponseWriter, request *http.Request) {
	// write database operations and logics here
	// if restAPI return json/xml

	// 1. read username and password from SPA
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	login_success := false
	//curl http://localhost:8080/user/login -X POST -d "mobile=15912465670&passwd=123456"
	if mobile == "15912465670" && passwd == "123456" {
		login_success = true
	}
	if !login_success {
		// set header and relevant responses
		Respond(writer, -1, nil, "Password or Username incorrect!")

	} else {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Respond(writer, 0, data, "")
	}

	//io.WriteString(writer, "Hello world!!") // test hello world
}

// define a struct to store the header information
type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Respond(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// define a struct and convert to json
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write([]byte(ret))
}

func main() {
	//bind the request and handler function
	http.HandleFunc("/user/login", user_login)

	//provide static resource support: File server
	//http.Handle("/", http.FileServer(http.Dir(".")))

	//provide designated directory
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}