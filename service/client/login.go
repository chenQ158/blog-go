package client
//
//import (
//	"blog-new/error"
//	"blog-new/models"
//	"context"
//	"github.com/smallnest/rpcx"
//	"log"
//)
//
//type LoginParam struct {
//	Username 	string	`json:"username"`
//	Password	string	`json:"password"`
//}
//
//type LoginReply struct {
//	ResultCode	string
//	ResultMsg	string
//	LoginUser	*models.User
//}
//
//func DoLogin(username, password string) (*models.User, error) {
//	var params = &LoginParam{Username:username, Password:password}
//	var reply *LoginReply
//	client := LoginSelector.Client
//	if client == nil {
//		client = rpcx.NewClient(LoginSelector)
//	}
//	err := client.Call(context.Background(), LoginN+".Dologin", &params, &reply)
//	if err != nil {
//		log.Fatal("rpc端访问错误")
//		return nil, errors.CommonError{ErrCode:"1001", ErrMsg:"访问错误"}
//	}
//	if reply.ResultCode != "" {
//		log.Fatal("rpc端访问错误")
//		return nil, errors.CommonError{ErrCode:"1002", ErrMsg:reply.ResultMsg}
//	}
//	if reply.LoginUser == nil {
//		log.Println("用户名或密码错误")
//		return nil, errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
//	}
//	return reply.LoginUser, nil
//}
