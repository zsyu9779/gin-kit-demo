package controllers

import (
	"gin-kit-demo/service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id")
	basicHandle := BasicController{Ctx: c}

	if valid.HasErrors() {
		basicHandle.Ok("failed")
		return
	}

	article:= service.GetArticleById(int64(id))


	basicHandle.Ok(article)
}
func GetArticleList(c *gin.Context) {

}
func AddArticle(c *gin.Context) {

}
func EditArticle(c *gin.Context) {

}
func DeleteArticle(c *gin.Context) {

}
func GenerateArticlePoster(c *gin.Context) {

}