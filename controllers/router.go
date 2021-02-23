package controllers

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	tagApi := e.Group("api/v1/tags")
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
	articleApi := e.Group("api/v1/articles")
	{
		//获取文章列表
		articleApi.GET("/list", GetArticles)
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


}