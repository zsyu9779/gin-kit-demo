package model

//easyjson:json
type ParentCommentResp struct {
	Comment
	ReplyComment []ReplyCommentResp `json:"replyComment"`
}
type ReplyCommentResp struct {
	Comment
	ParentComment Comment `json:"parentComment"`
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

func (c *Comment) GetCommentByBlogId() []Comment{
	var result []Comment
	Db.Where("blog_id = <> ? AND parent_comment_id = ?", c.BlogId,-1).Find(&c)
	return result
}


func (c *Comment) GetReplyCommentByCommentId() []Comment{
	var result []Comment
	Db.Where("parent_comment_id <> = ?",c.Id).Find(&result)
	return result
}
func (c *Comment) GetCommentById(id int64) Comment{
	var result  Comment
	Db.Find(&result,id)
	return result
}
func (c *Comment) AddComment() {
	Db.Create(&c)
}
