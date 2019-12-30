package bookctr

import (
	"github.com/kataras/iris/v12"
	"iriscli/datamodels"
	"iriscli/services/bookService"
	"log"
)

type BookController struct {
	Ctx 	iris.Context
	Service bookService.IBookService
}

func New() *BookController{
	return &BookController{
		Service: bookService.NewBookService(),
	}
}

// POST: /user/list
func (c *BookController) PostList()(result datamodels.Result)  {

	var m map[string]interface{}
	err := c.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}

	// 参数判断
	//if m["page"] == "" || m["page"] == nil {
	//	result.Code = -1
	//	result.Msg = "参数缺失 page"
	//	return
	//}

	return c.Service.GetBookList(m)
}