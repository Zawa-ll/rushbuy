package main

import (
	"context"

	"../backend/web/controllers"
	"../common"
	"../repositories"
	"../services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	tmplate := iris.HTML(". /backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)

	app.StaticWeb("/assets", ". /backend/web/assets")
	// Exception jump to the specified page
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	// Connecting to the database
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Error(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 5. Register Controller
	productRepository := repositories.NewProductManager("product", db)
	productSerivce := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productSerivce)
	product.Handle(new(controllers.ProductController))

	// 6. Starting services
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
