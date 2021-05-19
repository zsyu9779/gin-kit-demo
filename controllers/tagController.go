package controllers

import "github.com/gin-gonic/gin"

func GetTags(c *gin.Context) {
	//todo:page没传则设置默认值1
	basicHandle := BasicController{Ctx: c}
	basicHandle.Ok("")

}

func UpdateTag(c *gin.Context) {

}

func AddTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
