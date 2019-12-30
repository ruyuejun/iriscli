package userctr

import (
	"fmt"
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {

}

func New() *UserController{
	return &UserController{

	}
}

func (u *UserController) BeforeActivation(b mvc.BeforeActivation) {

	anyMiddleware := func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Inside /login")
		ctx.Next()
	}
	b.Handle("GET", "/login", "Login", anyMiddleware)
}

// GET: /user
func (u *UserController) Get() string {
	fmt.Println("get user...")
	return "Hey"
}

// POST: /user
func (u *UserController) Post() string {
	fmt.Println("post user...")
	return ""
}

// Put: /user
func (u *UserController) Put() string {
	fmt.Println("put user...")
	return ""
}

// GET: /user/login/{id:long}
func (u *UserController) Login() string {
	fmt.Println(u)
	return "MyCustomHandler says Hey"
}