package service

import (
	"blog-new/models"
	"blog-new/utils"
	"encoding/json"
	"fmt"
)

func AddComment(comment *models.Comment) error{
	db, err := utils.GetConn()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Create(comment).Error
	return err
}

func GetCommentsByTopicId(id int64) (*[]models.Comment, error) {
	db, err := utils.GetConn()
	if err != nil {
		return nil,err
	}
	defer db.Close()

	var comments []models.Comment
	err = db.Model(new(models.Comment)).Where("TopicId=?",id).Find(&comments).Error
	return &comments, err
}

// select topic.ID,topic.TITLE,topic.ATTACHMENT,topic.UPDATED,topic.VIEWS,topic.AUTHOR,category.TITLE,user.USERNAME from topic join category on category.ID = topic.CATEGORY_ID join user on user.ID = topic.USER_ID;
func GetTopicInfo() string {
	db, err := utils.GetConn()
	if err != nil {
		fmt.Println("数据库连接错误！")
		return ""
	}
	nativeSql := "select topic.ID,topic.TITLE,topic.ATTACHMENT,topic.UPDATED,topic.VIEWS,topic.AUTHOR,category.TITLE,user.USERNAME from topic join category on category.ID = topic.CATEGORY_ID join user on user.ID = topic.USER_ID"
	rows, _ := db.Raw(nativeSql).Rows()

	defer rows.Close()

	var ret []interface{}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				switch col.(type) {
				case []byte:
					record[columns[i]] = string(col.([]byte))
				default:
					record[columns[i]] = col
				}
			}
		}
		ret = append(ret, record)
	}

	str, _ := json.Marshal(ret)
	return string(str)
}