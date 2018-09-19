package service

import (
	"blog-new/models"
	"blog-new/utils"
	log "code.google.com/log4go"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/go-redis/redis"
	"time"
)

func GetRedisConn() (redis.Client, *error) {
	client, err := utils.GetRedis()

	if err != nil {
		log.Error("Redis数据库连接失败: " + err.Error())
		return *client,&err
	}
	return *client,nil
}

func GetRedisClusterConn() redis.ClusterClient {
	client := utils.REDIS_CLIENT
	return *client
}

func AddSession(uuid string, maxAge time.Duration, user *models.User, ctx *context.Context) *error {

	//client, _ := GetRedisConn()
	client := GetRedisClusterConn()

	userBs, err := json.Marshal(*user)
	if err != nil {
		return &err
	}

	// 将Token和User存入redis
	//client.Set(fmt.Sprintf("Token:%d", user.Id), uuid, maxAge)
	err = client.Set("Token:"+uuid, user.Id, maxAge).Err()
	if err != nil {
		fmt.Println("Set Token Error: ",err)
		return &err
	}
	err = client.Set(fmt.Sprintf("User:%d", user.Id), string(userBs), maxAge).Err()
	if err != nil {
		fmt.Println("Set Token Error: ",err)
		return &err
	}
	// 将Token存入cookie
	ctx.SetSecureCookie("chen", "Token", uuid, maxAge, "/")
	return nil
}

func GetSessionUser(token string) *models.User {
	//client, _ := GetRedisConn()
	client := GetRedisClusterConn()

	var user models.User
	userId,err := client.Get(fmt.Sprintf("Token:%s",token)).Result()
	if err!= nil {
		log.Error("Redis 查询错误：",err.Error())
		log.Error("Redis 查询Token不存在！")
		return nil
	}

	fmt.Println("userId:", userId)
	userStr,err := client.Get(fmt.Sprint("User:"+userId)).Result()
	if err != nil || len(userStr) == 0 {
		log.Error("Redis 查询错误：",err.Error())
		log.Error("Redis 查询User不存在！")
		return nil
	}

	fmt.Println("User:", user)
	err = json.Unmarshal([]byte(userStr), &user)
	if err != nil {
		log.Error("Redis 查询到User字符串转对象失败！")
		return nil
	}

	// fmt.Println(token, user)
	return &user
}

func DelSession(token string) *error {
	//client, _ := GetRedisConn()
	client := GetRedisClusterConn()

	userId := client.Get("Token:"+token)
	if len(userId.Val()) == 0 {
		return nil
	}
	client.Del("Token:"+token)
	userStr := client.Get("User:"+userId.Val())
	if len(userStr.Val()) == 0 {
		return nil
	}
	client.Del("User:"+userId.Val())
	// fmt.Println(token, userStr)
	return nil
}

