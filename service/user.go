package service

import (
	"IM-gossip-space/model"
	"IM-gossip-space/util"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

type UserService struct {
}

// user registeration
func (s *UserService) Register(mobile, plainpsswd, nickname, avatar, sex string) (user model.User, err error) {
	//检测手机号码是否存在,
	tmp := model.User{}
	_, err = DbEngine.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	//如果存在则返回提示已经注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}
	//否则拼接插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpsswd, tmp.Salt)
	tmp.Createat = time.Now()
	//token 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())
	//passwd =
	//md5 加密
	//返回新用户信息

	//插入 InserOne
	_, err = DbEngine.InsertOne(&tmp)
	//前端恶意插入特殊字符
	//数据库连接操作失败
	return tmp, err
}

func (s *UserService) Login(mobile, plainpsswd string) (user model.User, err error) {
	// search the user by mobile number
	//首先通过手机号查询用户
	tmp := model.User{}
	_, err = DbEngine.Where("mobile = ?", mobile).Get(&tmp)
	//如果没有找到
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	//查询到了比对密码
	if !util.ValidatePasswd(plainpsswd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	//刷新token,安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	//返回数据
	_, err = DbEngine.Where(" id = ?", tmp.Id).Cols("token").Update(&tmp)
	return tmp, err
}

// Find user by userID
func (s *UserService) Find(userId int64) (user model.User) {

	//首先通过手机号查询用户
	tmp := model.User{}
	DbEngine.Where("id = ?", userId).Get(&tmp)
	return tmp
}
