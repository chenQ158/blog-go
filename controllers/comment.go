package controllers

import (
	"blog-new/models"
	"blog-new/service"
	log "code.google.com/log4go"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type CommentController struct {
	beego.Controller
}

// @router /comment/add [post]
func (c *CommentController) Post() {
	var commentInfo models.CommentInfo
	err := c.ParseForm(&commentInfo)
	var comment models.Comment
	if err != nil {
		log.Error("评论：请求参数转换错误："+err.Error())
		c.Data["Error"] = "添加评论失败"
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.Data["PageName"] = "错误处理"
		c.TplName = "error.html"
		return
	} else {
		comment.Content = commentInfo.Content
		comment.Created = time.Now()
		comment.Name = commentInfo.Name
		comment.Email = commentInfo.Email
		comment.TopicId = commentInfo.TopicId
		comment.ParentId = commentInfo.ParentId
		err := service.AddComment(&comment)
		if err != nil {
			log.Error("评论：添加评论失败："+err.Error())
			c.Data["Error"] = "添加评论失败"
			c.Data["IsLogin"] = checkAccount(c.Ctx)
			c.Data["PageName"] = "错误处理"
			c.TplName = "error.html"
			return
		}
		err = service.UpdateTopicCommentCount(commentInfo.TopicId)
		if err != nil {
			log.Error("评论：更新评论数失败："+err.Error())
			c.Data["Error"] = "更新评论数失败"
			c.Data["IsLogin"] = checkAccount(c.Ctx)
			c.Data["PageName"] = "错误处理"
			c.TplName = "error.html"
			return
		}
	}
	c.Redirect(fmt.Sprintf("/user/topic/show?id=%d",commentInfo.TopicId), 302)
}
