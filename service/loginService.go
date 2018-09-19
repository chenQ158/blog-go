package service

import (
	"blog-new/models"
	"blog-new/utils"
	log "code.google.com/log4go"
)

func DoLogin(username, password string) (*models.User, error) {
	db, err := utils.GetConn()
	var user models.User
	if err != nil {
		log.Error("Mysql数据库连接错误！")
		return nil,err
	}
	defer db.Close()
	// select id,username,password from user where username=? and password=MD5(salt+?)
	db.Where("USERNAME=? and PASSWORD=MD5(SALT+?)", username, password).First(&user)
	return &user,nil
}


