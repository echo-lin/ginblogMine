package routes

import (
	v1 "ginblogMine/api/v1"
	"ginblogMine/middleware"
	"ginblogMine/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// user模块的路由接口
		auth.PUT("user/:id",v1.EditUser)
		auth.DELETE("user/:id",v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add",v1.AddCategory)

		auth.PUT("category/:id",v1.EditCate)
		auth.DELETE("category/:id",v1.DeleteCate)
		// 文章模块的的接口
		auth.POST("article/add",v1.AddArticle)


		auth.PUT("article/:id",v1.EditArt)
		auth.DELETE("article/:id",v1.DeleteArt)

		// 上传文件接口
		auth.POST("upload",v1.Upload)

	}

	router := r.Group("api/v1")
	{
		router.POST("user/add",v1.AddUser)
		router.GET("article",v1.GetArt)
		router.GET("category",v1.GetCate)
		router.GET("users",v1.GetUsers)
		router.GET("article/list/:id",v1.GetCateArtInfo)
		router.GET("article/info/:id",v1.GetArtInfo)
		router.POST("login",v1.Login)
	}

	r.Run(utils.HttpPort)
}
