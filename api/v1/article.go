package v1

import (
	"fmt"
	"ginblogMine/model"
	"ginblogMine/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context){
	// todo 添加用户
	var data model.Article

	_ = c.ShouldBindJSON(&data)
	fmt.Println(data,"哈哈哈")
	model.CreateArt(&data)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

// 查询文章
func GetArtInfo(c *gin.Context){
	var artInfo model.Article
	id,_ := strconv.Atoi(c.Param("id"))
	fmt.Println("ddddd",id)
	artInfo,code := model.GetArtInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":artInfo,
		"msg":errmsg.GetErrMsg(code),
	})
}

func GetCateArtInfo(c *gin.Context){
	cid,_ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	data,code := model.GetCateArtInfo(cid,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"msg":errmsg.GetErrMsg(code),
	})

}

// 查询用户列表
func GetArt(c *gin.Context){
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	data,code,total := model.GetArt(pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditArt(c *gin.Context){
	var data model.Article
	id,_ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	model.EditArt(id,&data)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})

}

// 删除用户
func DeleteArt(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}