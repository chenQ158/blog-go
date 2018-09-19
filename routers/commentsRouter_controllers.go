package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blog-new/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blog-new/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/user/category/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blog-new/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/user/category/del`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blog-new/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/user/category/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blog-new/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "TopicList",
			Router: `/user/category/topic_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:CommentController"] = append(beego.GlobalControllerRouter["blog-new/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/comment/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:ErrorController"] = append(beego.GlobalControllerRouter["blog-new/controllers:ErrorController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/error`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-new/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/comm/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-new/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/comm/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-new/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Reg",
			Router: `/comm/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-new/controllers:LoginController"],
		beego.ControllerComments{
			Method: "ToLogin",
			Router: `/comm/toLogin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-new/controllers:LoginController"],
		beego.ControllerComments{
			Method: "ToReg",
			Router: `/comm/toReg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-new/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-new/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:MainController"] = append(beego.GlobalControllerRouter["blog-new/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:TopicController"] = append(beego.GlobalControllerRouter["blog-new/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/user/topic/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:TopicController"] = append(beego.GlobalControllerRouter["blog-new/controllers:TopicController"],
		beego.ControllerComments{
			Method: "ToAdd",
			Router: `/user/topic/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:TopicController"] = append(beego.GlobalControllerRouter["blog-new/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/user/topic/del`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:TopicController"] = append(beego.GlobalControllerRouter["blog-new/controllers:TopicController"],
		beego.ControllerComments{
			Method: "ToEdit",
			Router: `/user/topic/edit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:TopicController"] = append(beego.GlobalControllerRouter["blog-new/controllers:TopicController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/user/topic/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:TopicController"] = append(beego.GlobalControllerRouter["blog-new/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Search",
			Router: `/user/topic/search`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:TopicController"] = append(beego.GlobalControllerRouter["blog-new/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `/user/topic/show`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-new/controllers:TopicController"] = append(beego.GlobalControllerRouter["blog-new/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/user/topic/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
