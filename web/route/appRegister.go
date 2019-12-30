package route

import (
	"github.com/kataras/iris/v12"
)

func RegisterRoute(app *iris.Application)  {

	// 注册最简单的示例路由 simple
	regSimpleRouter(app)

	// 注册完整MVC示例路由 book
	regBookRouter(app)

	// 注册用户路由
	regUserRouter(app)

}
