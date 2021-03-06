package controllers

import (
	"gin-kit-demo/service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

var articleService service.ArticleService
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id")
	basicHandle := BasicController{Ctx: c}

	if valid.HasErrors() {
		basicHandle.Ok("failed")
		return
	}

	article:= articleService.GetArticleRespById(int64(id))

	c.JSONP(200,article)
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