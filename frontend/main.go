package main

import (
	"context"

	"github.com/Zawa-ll/rushbuy/common"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"github.com/Zawa-ll/rushbuy/frontend/web/controllers"
	"github.com/Zawa-ll/rushbuy/repositories"
	"github.com/Zawa-ll/rushbuy/services"
	"github.com/kataras/iris/v12/sessions"

	"time"
)

func main() {

	app := iris.New()
	app.Logger().SetLevel("debug")
	tmplate := iris.HTML("./fronted/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	app.HandleDir("/public", "./frontend/web/public")
	app.HandleDir("/html", "./frontend/web/htmlProductShow")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "Error on the visited page!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	db, err := common.NewMysqlConn()
	if err != nil {

	}
	sess := sessions.New(sessions.Config{
		Cookie:  "AdminCookie",
		Expires: 600 * time.Minute,
	})
	// ctx controls the lifetime of requests or operations that should be canceled together.
	// ctx is useful for managing resources and controlling goroutines' execution in concurrent application
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	user := repositories.NewUserRepository("user", db)
	userService := services.NewService(user)
	userPro := mvc.New(app.Party("/user"))
	userPro.Register(userService, ctx, sess.Start)
	userPro.Handle(new(controllers.UserController))

	app.Run(
		iris.Addr("0.0.0.0:8082"),
		// iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
