package controllers

import (
	"gin-kit-demo/service/article_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetArticles(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	basicHandle := BasicController{Ctx: c}

	if valid.HasErrors() {
		basicHandle.Ok("failed")
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		basicHandle.Ok("failed")
		return
	}
	if !exists {
		basicHandle.Ok("failed")
		return
	}

	article, err := articleService.Get()
	if err != nil {
		basicHandle.Ok("failed")
		return
	}

	basicHandle.Ok(article)
}
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id")
	basicHandle := BasicController{Ctx: c}

	if valid.HasErrors() {
		basicHandle.Ok("failed")
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		basicHandle.Ok("failed")
		return
	}
	if !exists {
		basicHandle.Ok("failed")
		return
	}

	article, err := articleService.Get()
	if err != nil {
		basicHandle.Ok("failed")
		return
	}

	basicHandle.Ok(article)
}
func AddArticle(c *gin.Context) {

}
func EditArticle(c *gin.Context) {

}
func DeleteArticle(c *gin.Context) {

}
func GenerateArticlePoster(c *gin.Context) {

}