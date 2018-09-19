package controllers

import (
	"blog-new/service"
	log "code.google.com/log4go"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"time"
)

type LoginController struct {
	beego.Controller
}

// @router /comm/toLogin [get]
func (this *LoginController) ToLogin() {
	this.TplName = "login.html"
	this.Data["PageName"] = "登录"
	this.Data["IsHome"] = true
}

// @router /comm/logout [get]
func (this *LoginController) Logout() {
	token, exist := this.Ctx.GetSecureCookie("chen", "Token")
	if !exist {
		this.Redirect("/	", 302)//不能设置为301
		return
	}
	this.Ctx.SetCookie("Token", "", -1, "/")
	service.DelSession(token)
	this.Redirect("/	", 302)//不能设置为301
}

// @router /comm/login [post]
func (this *LoginController) Login() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autologin := this.Input().Get("autoLogin") == "on"

	if len(uname) == 0 || len(pwd) == 0 {
		log.Debug("用户名或密码为空！")
		// 返回页面
		this.TplName = "login.html"
		this.Data["Error"] = "用户名或密码为空！"
		return
	}

	user, err := service.DoLogin(uname, pwd)
	if err != nil {
		log.Error("数据库连接错误！")
		// 返回页面
		this.TplName = "login.html"
		this.Data["Error"] = "服务器内部错误，请联系管理员！"
		return
	}

	if user.Id == 0 {
		log.Debug("用户名或密码错误！")
		// 返回页面
		this.TplName = "login.html"
		this.Data["Error"] = "用户名或密码错误！"
		return
	}

	maxAge := 30 * time.Minute
	if autologin {
		maxAge = 1 << 32 - 1
	}

	// 登录成功后向session和cookie里添加用户信息
	sessionId := uuid.New().String()
	errptr := service.AddSession(sessionId, maxAge, user, this.Ctx)
	if errptr != nil {
		log.Error("登录失败！")
		// 返回页面
		this.TplName = "login.html"
		this.Data["Error"] = "登录失败，请联系管理员！"
		return
	}

	this.Redirect("/", 302)
}

// @router /login [get]
func (this *LoginController) Get() {

	exit := this.Input().Get("exit") == "true"
	//fmt.Println(exit)
	if !exit {
		this.TplName = "login.html"
		this.Data["PageName"] = "登录"
		this.Data["IsHome"] = true
	} else {
		token, exist := this.Ctx.GetSecureCookie("chen", "Token")
		if !exist {
			return
		}
		this.Ctx.SetCookie("Token", "", -1, "/")
		service.DelSession(token)
		this.Redirect("/	", 302)//不能设置为301
		return
	}
}

// @router /login [post]
func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autologin := this.Input().Get("autoLogin") == "on"

	if len(uname) == 0 || len(pwd) == 0 {
		log.Debug("用户名或密码为空！")
		// 返回页面
		this.TplName = "login.html"
		this.Data["Error"] = "用户名或密码为空！"
		return
	}

	user, err := service.DoLogin(uname, pwd)
	if err != nil {
		log.Error("数据库连接错误！")
		// 返回页面
		this.TplName = "login.html"
		this.Data["Error"] = "服务器内部错误，请联系管理员！"
		return
	}

	if user.Id == 0 {
		log.Debug("用户名或密码错误！")
		// 返回页面
		this.TplName = "login.html"
		this.Data["Error"] = "用户名或密码错误！"
		return
	}

	maxAge := 30 * time.Minute
	if autologin {
		maxAge = 1 << 32 - 1
	}

	// 登录成功后向session和cookie里添加用户信息
	sessionId := uuid.New().String()
	errptr := service.AddSession(sessionId, maxAge, user, this.Ctx)
	if errptr != nil {
		log.Error("登录失败！")
		// 返回页面
		this.TplName = "login.html"
		this.Data["Error"] = "登录失败，请联系管理员！"
		return
	}

	this.Redirect("/", 302)
}

// @router /comm/toReg [get]
func (this *LoginController) ToReg() {

}

// @router /comm/reg [post]
func (this *LoginController) Reg() {

}