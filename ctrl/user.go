package ctrl

import (
	"IM-gossip-space/model"
	"IM-gossip-space/service"
	"IM-gossip-space/util"
	"fmt"
	"math/rand"
	"net/http"
)

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	// write database operations and logics here
	// if restAPI return json/xml

	// 1. read username and password from SPA
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	//curl http://localhost:8080/user/login -X POST -d "mobile=15912465670&passwd=123456"
	user, err := userService.Login(mobile, passwd)
	//if mobile == "15912465670" && passwd == "123456" {
	//	login_success = true
	//}

	if err != nil {
		// set header and relevant responses
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}
}

var userService service.UserService

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainpasswd := request.PostForm.Get("passwd")
	nickName := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW
	user, err := userService.Register(mobile, plainpasswd, nickName, avatar, sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}

}
