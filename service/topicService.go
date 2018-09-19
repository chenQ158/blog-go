package service

import (
	"blog-new/models"
	"blog-new/utils"
	log "code.google.com/log4go"
	"encoding/json"
	"time"
)

func DelTopicUsingGorm(id, categoryId int64) (error) {
	db,err := utils.GetConn()
	if err != nil {
		return err
	}
	defer db.Close()
	//db.Where("id=?", id).Delete(models.Topic{})
	var topic = models.Topic{Id:id}
	db.Delete(&topic)
	err = db.Exec("UPDATE category SET TOPIC_COUNT=TOPIC_COUNT-1,TOPIC_TIME=sysdate() WHERE ID=?", categoryId).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTopicsUsingGorm(page, limit int64) (*[]models.TopicInfo, error) {
	db, err := utils.GetConn()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	//fmt.Println(page,)
	nativeSql := "select topic.ID ID,topic.TITLE TITLE,topic.ATTACHMENT ATTACHMENT,topic.REPLAY_COUNT REPLAY_COUNT,topic.UPDATED UPDATED,user.ID UID,user.USERNAME AUTHOR,category.ID CATEGORY_ID,category.TITLE CATEGORY_TITLE from topic left join category on topic.CATEGORY_ID=category.ID left join user on user.ID=topic.USER_ID limit ?, ?"
	resList := GeneralQuery(db, nativeSql, (page-1)*limit, limit)
	//fmt.Println("resList:", resList)
	btArr, err := json.Marshal(resList)
	// fmt.Println("GetTopicsByKeyword btArr:", string(btArr))
	if err != nil {
		log.Error("文章详情查询结果转json失败")
		return nil, err
	}
	var infoList []models.TopicInfo
	err = json.Unmarshal(btArr, &infoList)
	if err != nil {
		log.Error("json转文章详情失败")
		return nil, err
	}
	return &infoList, nil
}

func GetTopicByIdUsingGorm(id int64) (models.Topic, error) {
	db, err := utils.GetConn()
	var topic models.Topic
	if err != nil {
		return topic, err
	}
	defer db.Close()
	var cat models.Category
	err = db.First(&topic, "id=?", id).Error
	if err != nil {
		return topic,err
	}
	topic.Category = cat
	return topic, nil
}

func AddTopicUsingGorm(topic models.Topic) error {
	db, err := utils.GetConn()
	if err != nil {
		return err
	}
	defer db.Close()
	topic.Updated = time.Now()
	topic.Created = time.Now()
	topic.ReplayTime = time.Now()
	err = db.Create(&topic).Error
	if err != nil {
		return err
	}
	db.Exec("UPDATE category SET TOPIC_COUNT=TOPIC_COUNT+1,TOPIC_TIME=sysdate() WHERE ID=?", topic.CategoryId)
	return nil
}

func UpdateTopicUsingGorm(form models.TopicForm) error {
	db, err := utils.GetConn()
	if err != nil {
		return err
	}
	defer db.Close()
	var topic = models.Topic{Title:form.Title,Content:form.Content,Attachmemt:form.Attachment}
	err = db.Model(&topic).Where("Id=?", form.Id).Updates(topic).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTopicsByOffsetAndLimit(page, limit int64) (*[]models.Topic, error) {
	db, err := utils.GetConn()
	if err != nil {
		return nil,err
	}
	defer db.Close()
	var topics []models.Topic
	err = db.Limit(limit).Offset((page-1)*limit).Find(&topics).Error
	if err != nil {
		return &topics,err
	}
	return &topics,nil
}


func GetTopicsCount() int64 {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("数据库连接错误："+err.Error())
		return 0
	}
	defer db.Close()
	var count int64
	db.Model(&models.Topic{}).Count(&count)
	return count
}

func GetTopicsByCatId(catId, pageNum, limit int64) (*[]models.Topic, error) {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("数据库连接错误："+err.Error())
		return nil,err
	}
	defer db.Close()
	var topics []models.Topic
	err = db.Model(new(models.Topic)).Where("CATEGORY_ID=?", catId).Limit(limit).Offset((pageNum-1)*limit).Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return &topics,nil
}

func GetTopicsCountByCatId(catId int64) (int64, error) {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("数据库连接错误："+err.Error())
		return 0,err
	}
	defer db.Close()

	var count int64
	db.Model(new(models.Topic)).Where("CATEGORY_ID=?", catId).Count(&count)
	return count,nil
}

func UpdateTopicCommentCount(id int64) error {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("数据库连接错误："+err.Error())
		return err
	}
	defer db.Close()

	err = db.Exec("UPDATE topic SET REPLAY_COUNT = REPLAY_COUNT + 1 WHERE ID=?",id).Error
	return err
}


func GetTopicsByKeyword(keyword string) (*[]models.TopicInfo, error) {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("数据库连接错误："+err.Error())
		return nil,err
	}
	defer db.Close()

	nativeSql := "select topic.ID ID,topic.TITLE TITLE,topic.ATTACHMENT ATTACHMENT,topic.REPLAY_COUNT REPLAY_COUNT,topic.UPDATED UPDATED,user.ID UID,user.USERNAME AUTHOR,category.ID CATEGORY_ID,category.TITLE CATEGORY_TITLE from topic left join category on topic.CATEGORY_ID=category.ID left join user on user.ID=topic.USER_ID where topic.TITLE like ?"
	resList := GeneralQuery(db, nativeSql, keyword+"%")
	//fmt.Println("resList:", resList)
	btArr, err := json.Marshal(resList)
	//fmt.Println("GetTopicsByKeyword btArr:", string(btArr))
	if err != nil {
		log.Error("文章详情查询结果转json失败")
		return nil, err
	}
	var infoList []models.TopicInfo
	err = json.Unmarshal(btArr, &infoList)
	if err != nil {
		log.Error("json转文章详情失败")
		return nil, err
	}
	//fmt.Println("InfoList:", infoList)
	return &infoList, nil
}