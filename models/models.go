package models

import (
	"blog-new/utils"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	_DB_NAME		= "root@tcp(192.168.92.128:3306)/blog?charset=utf8"
	_MYSQL_DRIVER	= "mysql"
)

// 登录用户
type User struct {
	//主键
	Id			int64	`gorm:"column:ID;AUTO_INCREMENT;primary_key" json:"ID"`
	//登录名
	Username	string	`gorm:"column:USERNAME;unique" json:"USERNAME"`
	//密码
	Password	string	`gorm:"column:PASSWORD" json:"PASSWORD"`
	//盐
	Salt		string	`gorm:"column:SALT" json:"SALT"`
	//昵称
	Nickname	string	`gorm:"column:NICKNAME" json"NICKNAME"`
}

// 文章分裂
type Category struct {
	Id		int64		`gorm:"column:ID;AUTO_INCREMENT;primary_key" json:"ID"`						//主键
	Title	string		`gorm:"column:TITLE;unique" json:"TITLE"`					//分类名称
	Created	time.Time	`gorm:"column:CREATED" json:"CREATED"`			//创建时间
	TopicTime time.Time	`gorm:"column:TOPIC_TIME" json:"TOPIC_TIME"`		//更新时间
	TopicCount int64	`gorm:"column:TOPIC_COUNT" json:"TOPIC_COUNT"`	//分类下文章数量
	Topics	[]Topic
}

// 发表文章
type Topic struct {
	Id			int64		`gorm:"column:ID;AUTO_INCREMENT;primary_key" json:"ID"`			//主键
	Uid			int64		`gorm:"column:USER_ID" json:"USER_ID"`//用户id
	Title		string		`gorm:"column:TITLE" json:"TITLE"`		//文章标题
	Content		string		`gorm:"column:CONTENT;type:text" json:"CONTENT"`//文章内容
	Attachmemt	string		`gorm:"column:ATTACHMENT" json:"ATTACHMENT"`//标签
	Created		time.Time	`gorm:"column:CREATED" json"CREATED"`	//创建时间
	Updated		time.Time	`gorm:"column:UPDATED" json"UPDATED"`	//更新时间
	Views		int64		`gorm:"column:VIEWS" json"VIEWS"`		//浏览数
	Author		string		`gorm:"column:AUTHOR" json"AUTHOR"`	//作者
	ReplayTime	time.Time	`gorm:"column:REPLAY_TIME" json"REPLAY_TIME"`	//最后评论时间
	ReplayCount	int64		`gorm:"column:REPLAY_COUNT" json"REPLAY_COUNT"` 	//评论数
	ReplayLastUserId int64	`gorm:"column:REPLAY_LAST_USER_ID" json"REPLAY_LAST_USER_ID"`	//最后评论用户Id
	CategoryId	int64		`gorm:"column:CATEGORY_ID;index" json:"CATEGORY_ID"`	//所属分类
	Category 	Category
}

// 文章概述
type TopicSummary struct {
	// 主键与topic主键相同
	Id			int64		`gorm:"column:ID" json:"ID"`
	// 作者
	Author		string		`gorm:"column:AUTHOR" json"AUTHOR"`
	// 概述内容
	Content		string		`gorm:"column:CONTENT" json:"CONTENT"`
	// 评论数
	CommentCount int64		`gorm:"column:COMMENT_COUNT" json:"COMMENT_COUNT"`
	// 浏览数
	Views		int64		`gorm:"column:VIEWS" json:"VIEWS"`
	// 创建日期
	Created		time.Time	`gorm:"column:CREATED" json:"CREATED"`
}

// 事件实体
type EventEntity struct {
	// 主键
	Id			int64		`gorm:"column:ID;AUTO_INCREMENT;primary_key"		json:"ID"`
	// 被评论或点赞主体类型：默认为文章0
	EntityType	int8		`gorm:"column:ENTITY_TYPE" json:"ENTITY_TYPE"`
	// 被评论或点赞主题Id
	EntityId	int64		`gorm:"column:ENTITY_ID" json:"ENTITY_ID"`
	// 评论或点赞人ID
	ActorId		int64		`gorm:"column:ACTOR_ID" json:"ACTOR_ID"`
	// 被评论或点赞主体所属人
	ReceiverID	int64		`gorm:"column:RECEIVER_ID" json:"RECEIVER_ID"`
	// 评论内容，如果是点赞则为空
	Content		string		`gorm:"column:CONTENT" json:"CONTENT"`
	// 评论或点赞时间
	Created		time.Time	`gorm:"column:CREATED" json:"CREATED"`
	// 父评论ID 如果没有则为0，如果是点赞则为-1
	ParentId	int64		`gorm:"column:PARENT_ID" json:"PARENT_ID"`
}

// 评论
type Comment struct {
	// 主键
	Id			int64		`gorm:"column:ID;AUTO_INCREMENT;primary_key"  json:"ID"`
	// 文章Id
	TopicId		int64		`gorm:"column:TOPICID"		json:"TOPICID"`
	// 评论内容
	Content		string		`gorm:"column:CONTENT"		json:"CONTENT"`
	// 评论人邮箱
	Email		string		`gorm:"column:EMAIL"			json:"EMAIL"`
	// 评论人
	Name		string		`gorm:"column:NAME"			json:"NAME"`
	// 父评论，如果没有则为0
	ParentId	int64		`gorm:"column:PARENT_ID"	json:"PARENT_ID"`
	// 创建时间
	Created		time.Time	`gorm:"column:CREATED"		json:"CREATED"`
}

//func (*Category) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("Id", 123)
//	return nil
//}

func RegisterDBUsingBeego() {
	orm.RegisterDriver(_DB_NAME, orm.DRMySQL)
	orm.RegisterModel(new(User), new(Category), new(Topic), new(EventEntity))
	orm.RegisterDataBase("default", _MYSQL_DRIVER, _DB_NAME, 10)
	orm.RunSyncdb("default", false, true)
}

func RegisterDBUsingGorm() {
	db,err := utils.GetConn()
	if err != nil {
		// panic
		return
	}

	// 禁用复数建表
	db.SingularTable(true)
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("charset", "utf8")
	// 自动迁移模式
	db.AutoMigrate(&User{}, &Category{}, &Topic{}, &EventEntity{}, &Comment{})
}

// 文章提交表单实体
type TopicForm struct {
	Id				int64		`form:"id"`
	Title 			string		`form:"title"`
	CatId			int64		`form:"category"`
	Content			string		`form:"content"`
	Attachment		string		`form:"attachment"`
	Author			string		`form:"author"`
}

// 评论请求表单实体
type CommentInfo struct {
	TopicId			int64		`form:"topicId"`
	Content			string		`form:"content"`
	Name			string		`form:"name"`
	Email			string		`form:"email"`
	ParentId		int64		`form:"parentId"`
}

// 文章详情，用于页面展示
type TopicInfo struct {
	Id			int64		`json:"ID"`
	Title		string		`json:"TITLE"`		//文章标题
	Attachmemt	string		`json:"ATTACHMENT"`//标签
	Updated		time.Time	`json:"UPDATED"`	//更新时间
	ReplayCount	int64		`json:"REPLAY_COUNT"`		//评论数
	CategoryId	int64		`json:"CATEGORY_ID"`
	CategoryTitle  string	`json:"CATEGORY_TITLE"`
	Uid			int64		`json:"UID"`
	Author		string		`json:"AUTHOR"`
}

