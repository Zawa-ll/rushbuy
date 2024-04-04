package controllers

import (
	"strconv"

	"github.com/Zawa-ll/rushbuy/datamodels"
	"github.com/Zawa-ll/rushbuy/services"
	"github.com/Zawa-ll/rushbuy/tool"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type UserController struct {
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

func (c *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (c *UserController) PostRegister() {
	var (
		nickName = c.Ctx.FormValue("nickName")
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	//ozzo-validation
	user := &datamodels.User{
		UserName:     userName,
		NickName:     nickName,
		HashPassword: password,
	}

	_, err := c.Service.AddUser(user)
	c.Ctx.Application().Logger().Debug(err)
	if err != nil {
		c.Ctx.Redirect("/user/error")
		return
	}
	c.Ctx.Redirect("/user/login")
	return
}

func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (c *UserController) PostLogin() mvc.Response {
	// 1.Get information about the form submitted by the user
	var (
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	// 2.Verify that the account password is correct
	user, isOk := c.Service.IsPwdSuccess(userName, password)
	if !isOk {
		return mvc.Response{
			Path: "/user/login",
		}
	}

	// 3.Write the user ID to the cookie
	tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))
	c.Session.Set("userID", strconv.FormatInt(user.ID, 10))

	return mvc.Response{
		Path: "/product/",
	}

}
