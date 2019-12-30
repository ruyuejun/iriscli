
package userService

import (
	"iriscli/middleware"
	"iriscli/datamodels"
	userRepo2 "iriscli/datarepo/userRepo"
	"iriscli/utils"
	// "fmt"
	// "github.com/spf13/cast"
	// "log"
)


type userServices struct {

}

func NewUserServices() IUser {
	return &userServices{}
}

var userRepo = userRepo2.NewUserRepo()

/*
登录
*/
func (u userServices) Login(m map[string]string) (result datamodels.Result) {

	if m["username"] == "" {
		result.Code = -1
		result.Msg = "请输入用户名！"
		return
	}
	if m["password"] == "" {
		result.Code = -1
		result.Msg = "请输入密码！"
		return
	}
	user := userRepo.GetUserByUserNameAndPwd(m["username"],utils.GetMD5String(m["password"]))
	if user.ID == 0 {
		result.Code = -1
		result.Msg = "用户名或密码错误!"
		return
	}
	user.Token = middleware.GenerateToken(user)
	result.Code = 0
	result.Data = user
	result.Msg = "登录成功"
	return
}

/*
保存
*/
func (u userServices) Save(user datamodels.User) (result datamodels.Result){
	//添加
	if user.ID == 0 {
		agen := userRepo.GetUserByUsername(user.Username)
		if agen.ID != 0 {
			result.Msg = "登录名重复,保存失败"
			return
		}
	}
	code,p := userRepo.Save(user)
	if code == -1 {
		result.Code = -1
		result.Msg = "保存失败"
		return
	}
	result.Code = 0
	result.Data = p
	return
}