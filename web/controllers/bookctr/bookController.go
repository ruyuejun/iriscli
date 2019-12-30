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

func NewBookController() *BookController{
	return &BookController{
		Service: bookService.NewBookService(),
	}
}

func (g *BookController) PostList()(result datamodels.Result)  {

	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}

	// 参数判断
	//if m["page"] == "" || m["page"] == nil {
	//	result.Code = -1
	//	result.Msg = "参数缺失 page"
	//	return
	//}

	return g.Service.GetBookList(m)
}