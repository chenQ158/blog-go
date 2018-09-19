package controllers

import (
	"blog-new/constant"
	"blog-new/service"
	"blog-new/utils"
	log "code.google.com/log4go"
	"github.com/astaxie/beego"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

// @router /user/category/add [get]
func (this *CategoryController) Add() {
	name := this.Input().Get("name")
	if len(name) == 0 {
		this.Data["Error"] = constant.VARIABLE_EXCHANGE_ERROR
		this.Redirect("/error", 302)
		return
	}
	res, err := service.AddCatUsingGorm(name)

	//fmt.Println(res)
	if err != nil || res == false {
		this.Data["Error"] = constant.CATEGORY_ADD_ERROR
		this.Redirect("/error", 302)
		return
	}
	this.Redirect("/user/category/list", 302)
}

// @router /user/category/del [get]
func (this *CategoryController) Del() {
	idStr := this.Input().Get("id")
	if len(idStr) == 0 {
		this.Data["Error"] = constant.VARIABLE_EXCHANGE_ERROR
		this.Redirect("/error", 302)
		return
	}
	id,err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		this.Data["Error"] = constant.VARIABLE_EXCHANGE_ERROR
		this.Redirect("/error", 302)
		return
	}

	res,err := service.DelCatUsingGorm(id)
	if err != nil || res == false {
		this.Data["Error"] = constant.CATEGORY_DEL_ERROR
		this.Redirect("/error", 302)
		return
	}

	this.Redirect("/user/category/list", 302)
}

// @router /user/category/topic_list [get]
func (this *CategoryController) TopicList() {
	idStr := this.Input().Get("id")
	if len(idStr) == 0 {
		this.Data["Error"] = constant.VARIABLE_EXCHANGE_ERROR
		this.Redirect("/error", 302)
		return
	}
	id,err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		this.Data["Error"] = constant.VARIABLE_EXCHANGE_ERROR
		this.Redirect("/error", 302)
		return
	}

	pageInput := this.Input().Get("page")
	limitInput := this.Input().Get("limit")
	var pageNum, limit int64

	if len(pageInput) != 0 {
		pageNum,err = strconv.ParseInt(pageInput, 10, 64)
		if err != nil {
			pageNum = 1
			log.Error("参数类型转换错误： "+err.Error())
		}
	}
	limit = 5
	if len(limitInput) != 0 {
		limit, err = strconv.ParseInt(limitInput, 10, 64)
		if err != nil {
			log.Error("参数类型转换错误： "+err.Error())
		}
	}

	topics,err := service.GetTopicsByCatId(id, pageNum, limit)
	for i:=0; i<len(*topics); i++ {
		(*topics)[i].Content = utils.VerifyHtml((*topics)[i].Content)
		if len((*topics)[i].Content) > 1000 {
			(*topics)[i].Content = (*topics)[i].Content[:1000]
		}
	}

	if err != nil {
		this.Data["Error"] = constant.VARIABLE_SEARCH_ERROR
		this.Redirect("/error", 302)
		return
	}
	count,err := service.GetTopicsCountByCatId(id)
	if err != nil {
		this.Data["Error"] = constant.VARIABLE_SEARCH_ERROR
		this.Redirect("/error", 302)
		return
	}
	page := utils.PageUtil(count, pageNum, limit, *topics)

	if len(*topics) == 0{
		this.Data["IsNull"] = true
	} else {
		this.Data["IsNull"] = false
	}

	this.Data["Page"] = page
	this.Data["CatId"] = id
	this.Data["PageName"] = "分类文章"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsCategory"] = true
	this.TplName = "topic_in_cat.html"
}

// @router /user/category/list [get]
func (this *CategoryController) List() {
	pageStr := this.Input().Get("page")
	var page = 1
	var err error
	if len(pageStr) != 0 {
		page,err = strconv.Atoi(pageStr)
		if err != nil {
			log.Error(err)
			this.Data["Error"] = constant.VARIABLE_SEARCH_ERROR
			this.Redirect("/error", 302)
			return
		}
	}
	page64,_ := strconv.ParseInt(strconv.Itoa(page), 10, 64)

	cats,err :=  service.GetCatsUsingGorm(page64, 12)
	count,_ := service.GetCatCount()
	pageList := utils.PageUtil(count, page64, 12, cats)

	this.Data["Page"] = pageList
	this.Data["Categories"] = cats
	this.Data["PageName"] = "分类"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
}


