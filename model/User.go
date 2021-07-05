package model

import (
	"encoding/base64"
	"ginblogMine/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(20);not null" json:"Password" validate:"required,min=4,max=12"`
	Role int `gorm:"type:int;DEFAULT:2" json:"Role" validate:"required,gte=2"`
}

// 查询用户是否存在
func CheckUser(name string)(code int){
	var user User
	db.Select("id").Where("username = ?" ,name).First(&user)
	if user.ID > 0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 编辑用户
func EditUser(id int, data *User) int{
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Update(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 创建用户
func CreateUser(data *User)int{

	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int)int{
	var user User
	err := db.Where(" id = ?", id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 加密密码
func (u *User)BeforeSave(){
	u.Password = ScryptPw(u.Password)
}

// 获取用户
func GetUsers(pageSize int, pageNum int)[]User{
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return users
}

// 加密密码
func ScryptPw(password string)string{
	const KeyLen = 10
	salt := make([]byte,8)
	salt = []byte{12,32,4,6,66,22,222,11}
	HashPw,err := scrypt.Key([]byte(password),salt,16384,8,1,KeyLen)
	if err != nil{
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

func CheckLogin(username string, password string)int{
	var user User
	db.Where(" username = ? ", username).Find(&user)
	if user.ID == 0{
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password{
		return errmsg.ERROR_PASSWARD_WRONG
	}
	if user.Role != 1{
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}









