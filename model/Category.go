package model

import (
	"ginblogMine/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(10);not null" json:"name"`
}


// 查询分类是否存在
func CheckCategory(name string)(code int){
	var cate Category
	db.Select("id").Where("name = ?" ,name).First(&cate)
	if cate.ID > 0{
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

// 编辑用户
func EditCate(id int, data *Category) int{
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Update(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 创建用户
func CreateCate(data *Category)int{

	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteCate(id int)int{
	var cate Category
	err := db.Where(" id = ?", id).Delete(&cate).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 获取用户
func GetCates(pageSize int, pageNum int)([]Category,int){
	var cate []Category
	var total int
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,0
	}
	return cate,total
}