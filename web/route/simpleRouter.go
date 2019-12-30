package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iriscli/web/controllers/simplectr"
)

func regSimpleRouter(app *iris.Application){
	mvc.New(app.Party("/simple")).Handle(new(simplectr.SimpleController))
}