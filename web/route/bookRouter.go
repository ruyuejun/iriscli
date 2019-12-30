package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iriscli/web/controllers/bookctr"
)

func regBookRouter(app *iris.Application) {
	mvc.New(app.Party("/book")).Handle(bookctr.NewBookController())
}