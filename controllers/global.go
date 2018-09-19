package controllers

import (
	"blog-new/service"
	"github.com/astaxie/beego/context"
)

func checkAccount(ctx *context.Context) bool {
	// 判断session中是否有Token
	//token, ok := ctx.Input.Session("Token").(string)
	//fmt.Println(token, ok)
	//if !ok || len(token) == 0 {
	//	return false
	//}
	// 判断session中是否有CurrentUser
	//user , ok := ctx.Input.Session("CurrentUser").(models.User)
	//fmt.Println(user, ok)
	//if !ok || user.Id == 0 {
	//	return false
	//}

	token, exist := ctx.GetSecureCookie("chen", "Token")
	if !exist {
		return false
	}
	userptr := service.GetSessionUser(token)
	if userptr == nil {
		return false
	}

	return true
}

