package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iriscli/web/controllers/userctr"
)


func regUserRouter(app *iris.Application) {
	mvc.Configure(app.Party("/user"), func(app *mvc.Application) {
		app.Handle(new(userctr.UserController))
	})
}
