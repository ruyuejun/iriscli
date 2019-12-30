package simplectr

import (
	"github.com/kataras/iris/v12/mvc"
)

type SimpleController struct{}

func New() *SimpleController{
	return &SimpleController{
	}
}

func (c *SimpleController) Get() mvc.Result {
	return mvc.Response{
		Err:  nil,
		Code: 400,
		Text: "simple...",
	}
}

func (c *SimpleController) GetBy(msg string) mvc.Result {
	return mvc.Response{
		Err: nil,
		Code: 400,
		Text: msg,
	}
}