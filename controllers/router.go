package controllers

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	tagApi := e.Group("api/v1/tag")
	{
		//获取标签列表
		tagApi.GET("",Wrapper(GetTags))
		//新建标签
		tagApi.POST("", Wrapper(AddTag))
		//更新指定标签
		tagApi.PUT("/:id", Wrapper(UpdateTag))
		//删除指定标签
		tagApi.DELETE("/:id", Wrapper(DeleteTag))
	}
	typeApi := e.Group("api/v1/type")
	{
		//获取分类列表
		typeApi.GET("",Wrapper(GetTypes))
		//新建分类
		typeApi.POST("",Wrapper(AddType))
		//更新指定分类
		typeApi.PUT("/:id",Wrapper(UpdateType))
		//删除指定分类
		typeApi.DELETE("/:id",Wrapper(DeleteType))
	}
	userApi := e.Group("api/v1/user")
	{
		//登录
		userApi.POST("/login",Wrapper(Login))
		//获取用户列表
		userApi.GET("/",Wrapper(GetUsers))
		//新建用户
		userApi.POST("",Wrapper(AddUser))
		//修改用户
		userApi.PUT("/:id",Wrapper(UpdateUser))
		//删除指定用户
		userApi.DELETE("/:id",Wrapper(DeleteUser))
	}
	articleApi := e.Group("api/v1/articles")
	{
		//获取文章列表
		articleApi.GET("/list", Wrapper(GetArticleList))
		//获取指定文章
		articleApi.GET("", Wrapper(GetArticle))
		//新建文章
		articleApi.POST("", Wrapper(AddArticle))
		//更新指定文章
		articleApi.PUT("", Wrapper(EditArticle))
		//删除指定文章
		articleApi.DELETE("", Wrapper(DeleteArticle))
		//生成文章海报
		articleApi.POST("/poster/generate", Wrapper(GenerateArticlePoster))
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
