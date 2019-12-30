package simplectr

import (
	"fmt"
	"github.com/kataras/iris/v12/mvc"
)

type SimpleController struct{}

func (c *SimpleController) Get() mvc.Result {
	fmt.Println("simple...")
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