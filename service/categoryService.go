package service

import (
	"blog-new/models"
	"blog-new/utils"
	"time"
)

func AddCatUsingGorm(name string) (bool, error){
	db,err := utils.GetConn()
	if err != nil {
		return false, nil
	}
	defer db.Close()
	cat := models.Category{Title:name, Created:time.Now(), TopicTime:time.Now()}
	err = db.Create(&cat).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetCatsUsingGorm(page,limit int64) (*[]models.Category, error) {
	db, err := utils.GetConn()
	var cats []models.Category
	if err != nil {
		return &cats, err
	}
	defer db.Close()
	err = db.Model(new(models.Category)).Limit(limit).Offset((page-1)*limit).Find(&cats).Error
	if err != nil {
		return &cats, err
	}
	return &cats, nil
}

func GetAllCats() (*[]models.Category, error) {
	db, err := utils.GetConn()
	var cats []models.Category
	if err != nil {
		return &cats, err
	}
	defer db.Close()

	err = db.Model(new(models.Category)).Find(&cats).Error
	return &cats, err
}

func GetCatCount() (int64, error) {
	db, err := utils.GetConn()
	if err != nil {
		return 0, nil
	}
	defer db.Close()

	var count int64
	err = db.Model(new(models.Category)).Count(&count).Error
	return count, err
}

func DelCatUsingGorm(id int64) (bool, error) {
	db, err := utils.GetConn()
	if err != nil {
		return false, nil
	}
	defer db.Close()
	err = db.Delete(models.Category{Id:id}).Error
	if err != nil {
		return false, err
	}
	return true,nil
}

func GetCatsByKeyUsingGorm(keyword string) (*[]models.Category,error) {
	db, err := utils.GetConn()
	var cats []models.Category
	if err != nil {
		return nil,err
	}
	defer db.Close()
	err = db.Where("TITLE like ?", keyword+"%").Find(&cats).Error
	if err != nil {
		return nil, err
	}
	return &cats, nil
}

