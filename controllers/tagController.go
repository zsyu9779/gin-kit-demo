package controllers

import (
	"gin-kit-demo/model"
	"gin-kit-demo/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var tagService service.TagService
func GetTags(c *gin.Context) {
	//todo:page没传则设置默认值1
	pageSize, _ := strconv.Atoi(c.Query("size"))
	pageIndex, _ := strconv.Atoi(c.Query("page"))
	//basicHandle := BasicController{Ctx: c}
	c.JSONP(200,tagService.GetTags(pageSize,pageIndex))

}

func UpdateTag(c *gin.Context) {

}

func AddTag(c *gin.Context) {
	name := c.Request.FormValue("name")
	color := c.Request.FormValue("color")
	tag :=model.Tag{
		Name:  name,
		Color: color,
	}
	tagService.AddTag(&tag)
	c.JSON(200,0)
}

func DeleteTag(c *gin.Context) {

}
