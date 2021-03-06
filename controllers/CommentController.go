package controllers

import (
	"gin-kit-demo/model"
	"gin-kit-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

var commentService service.CommentService

func GetComments(c *gin.Context) {
	id, e := strconv.Atoi(c.Query("id"))
	if e != nil {
		logrus.Error(e)
		logrus.Info(c.Param("id"))
	}
	result := commentService.GetComments(int64(id))
	//str, err := json.Marshal(result)
	//if err != nil {
	//	logrus.Error(err)
	//}
	c.JSONP(200, result)
}

func AddComment(c *gin.Context) {
	parentCommentId, _ := strconv.Atoi(c.Request.FormValue("parentComment.id"))
	blogId, _ := strconv.Atoi(c.Request.FormValue("blog.id"))
	nickName := c.Request.FormValue("nickname")
	email := c.Request.FormValue("email")
	content := c.Request.FormValue("content")
	avatar := c.Request.FormValue("avatar")

	comment := model.Comment{
		Email:           email,
		Avatar:          avatar,
		AdminComment:    false,
		NickName:        nickName,
		Content:         content,
		BlogId:          int64(blogId),
		ParentCommentId: int64(parentCommentId),
	}
	commentService.AddComment(comment)
	c.JSON(200, 0)

}
