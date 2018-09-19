package routers

import (
	"blog-new/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.MainController{},&controllers.LoginController{},&controllers.CategoryController{},&controllers.TopicController{},&controllers.CommentController{},&controllers.ErrorController{})
   //beego.Router("/", &controllers.MainController{})
	////beego.Router("/home", &controllers.MainController{})
	//beego.Router("/login", &controllers.LoginController{})
	//beego.Router("/user/category", &controllers.CategoryController{})
	//beego.Router("/user/topic", &controllers.TopicController{})
   //beego.Router("/comment", &controllers.CommentController{})
   //beego.Router("/error", &controllers.ErrorController{})
}
