package controllers

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	tagApi := e.Group("api/v1/tag")
	{
		//获取标签列表
		tagApi.GET("", GetTags)
		//新建标签
		tagApi.POST("", AddTag)
		//更新指定标签
		tagApi.PUT("/:id", UpdateTag)
		//删除指定标签
		tagApi.DELETE("/:id", DeleteTag)
	}
	typeApi := e.Group("api/v1/type")
	{
		//获取分类列表
		typeApi.GET("")
		//新建分类
		typeApi.POST("")
		//更新指定分类
		typeApi.PUT("/:id")
		//删除指定分类
		typeApi.DELETE("/:id")
	}
	userApi := e.Group("api/v1/user")
	{
		//登录
		userApi.POST("/login")
		//获取用户列表
		userApi.GET("/")
		//新建用户
		userApi.POST("")
		//修改用户
		userApi.PUT("/:id")
		//删除指定用户
		userApi.DELETE("/:id")
	}
	articleApi := e.Group("api/v1/articles")
	{
		//获取文章列表
		articleApi.GET("/list", GetArticleList)
		//获取指定文章
		articleApi.GET("", GetArticle)
		//新建文章
		articleApi.POST("", AddArticle)
		//更新指定文章
		articleApi.PUT("", EditArticle)
		//删除指定文章
		articleApi.DELETE("", DeleteArticle)
		//生成文章海报
		articleApi.POST("/poster/generate", GenerateArticlePoster)
	}
	archiveApi := e.Group("api/v1/archive")
	{
		archiveApi.GET("map")
		archiveApi.GET("count")
	}
	commentApi := e.Group("api/v1/comment")
	{
		commentApi.GET("")
		commentApi.POST("")
	}

}
