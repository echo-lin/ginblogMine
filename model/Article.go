package model

import (
	"fmt"
	"ginblogMine/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category	`gorm:"foreignkey:Cid"`
	gorm.Model
	Title string	`gorm:"type:varchar(100);not null" json:"title"`
	Cid int	`gorm:"type:int;not null" json:"cid"`
	Desc string	`gorm:"type:varchar(200)" json:"desc"`
	Content string	`gorm:"type:longtest" json:"content"`
	Img string	`gorm:"type:varchar(100)" json:"img"`
}


// 编辑用户
func EditArt(id int, data *Article) int{
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&art).Where("id = ?", id).Update(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 创建用户
func CreateArt(data *Article)int{

	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteArt(id int)int{
	var art Article
	err := db.Where(" id = ?", id).Delete(&art).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 获取文章列表
func GetArt(pageSize int, pageNum int)([]Article,int){
	var artticleList []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&artticleList).Error
	fmt.Println(pageSize,pageNum,err)
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,errmsg.ERROR
	}
	return artticleList,errmsg.SUCCSE
}

// 查询单个文章
func GetArtInfo(id int)(Article,int){
	var artInfo Article
	err := db.Preload("Category").Where("id = ?",id).First(&artInfo).Error
	if err != nil{
		return artInfo,errmsg.ERROR_ART_NOT_EXIST
	}
	return artInfo,errmsg.SUCCSE
}

// 查询分类下的文章
func GetCateArtInfo(id int,pageSize int, pageNum int)(Article,int){
	var art Article
	err := db.Preload("Category").Where("cid = ?",id).Limit(pageSize).Offset((pageNum - 1)*pageSize).Find(&art).Error
	if err != nil{
		return art,errmsg.ERROR_ART_NOT_EXIST
	}

	return art,errmsg.SUCCSE

}