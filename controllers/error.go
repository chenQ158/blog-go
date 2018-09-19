package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

// @router /error [get]
func (e *ErrorController) Get() {
	e.TplName = "error.html"
	err := e.Input().Get("Error")
	e.Data["Error"] = err
}
