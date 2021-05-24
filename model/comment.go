package model

import "github.com/sirupsen/logrus"

//easyjson:json
type ParentCommentResp struct {
	Id              int64              `json:"id"`
	Email           string             `json:"email"`
	Avatar          string             `json:"avatar"`
	AdminComment    bool               `json:"admin_comment"`
	NickName        string             `json:"nickname"`
	Content         string             `json:"content"`
	BlogId          int64              `json:"blog_id"`
	ParentCommentId int64              `json:"parent_comment_id"`
	Ctime           int64              `json:"create_time"`
	ReplyComment    []ReplyCommentResp `json:"replyComment"`
}
type ReplyCommentResp struct {
	Id              int64   `json:"id"`
	Email           string  `json:"email"`
	Avatar          string  `json:"avatar"`
	AdminComment    bool    `json:"admin_comment"`
	NickName        string  `json:"nickname"`
	Content         string  `json:"content"`
	BlogId          int64   `json:"blog_id"`
	ParentCommentId int64   `json:"parent_comment_id"`
	Ctime           int64   `json:"create_time"`
	ParentComment   Comment `json:"parentComment"`
}

//easyjson:json
type Comment struct {
	Id              int64  `json:"id"`
	Email           string `json:"email"`
	Avatar          string `json:"avatar"`
	AdminComment    bool   `json:"admin_comment"`
	NickName        string `json:"nickname"`
	Content         string `json:"content"`
	BlogId          int64  `json:"blog_id"`
	ParentCommentId int64  `json:"parent_comment_id"`
	Ctime           int64  `json:"create_time"`
}

func (c *Comment) GetCommentByBlogId() []Comment {
	var result []Comment
	err := Db.Where("blog_id = ? AND parent_comment_id = ?", c.BlogId, -1).Find(&result).Error
	if err != nil {
		logrus.Error(err)
	}
	return result
}

func (c *Comment) GetReplyCommentByCommentId() []Comment {
	var result []Comment
	Db.Where("parent_comment_id = ?", c.Id).Find(&result)
	return result
}
func (c *Comment) GetCommentById(id int64) Comment {
	var result Comment
	Db.Find(&result, id)
	return result
}
func (c *Comment) AddComment() {
	Db.Create(&c)
}
