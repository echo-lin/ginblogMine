package v1

import (
	"fmt"
	"ginblogMine/model"
	"ginblogMine/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加用户
func AddCategory(c *gin.Context){
	// todo 添加用户
	var data model.Category

	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE{
		model.CreateCate(&data)
	}

	if code == errmsg.ERROR_CATENAME_USED{
		code = errmsg.ERROR_CATENAME_USED
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetCate(c *gin.Context){
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	fmt.Println(pageNum,pageSize)
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	data,total := model.GetCates(pageSize,pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditCate(c *gin.Context){
	var data model.Category
	id,_ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE{
		model.EditCate(id,&data)
	}

	if code == errmsg.ERROR_CATENAME_USED{
		c.Abort()
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})

}

// 删除用户
func DeleteCate(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}