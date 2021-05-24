package service

import (
	"encoding/json"
	"gin-kit-demo/model"
)

type CommentService struct {

}

func (c *CommentService) AddComment(comment model.Comment) {
	comment.AddComment()
}
func (c *CommentService) GetComments(blogId int64) []model.ParentCommentResp {
	var result []model.ParentCommentResp
	comment := model.Comment{BlogId: blogId}
	commentList:=comment.GetCommentByBlogId()
	for _, item := range commentList {
		var commentWithReply model.ParentCommentResp
		str, _ := json.Marshal(item)
		json.Unmarshal(str,&commentWithReply)
		replyComment := c.HandleReplyComments(item.GetReplyCommentByCommentId())
		commentWithReply.ReplyComment = replyComment
		result = append(result, commentWithReply)
	}
	return result
}

func (c *CommentService) HandleReplyComments(list []model.Comment) []model.ReplyCommentResp {
	var result []model.ReplyCommentResp
	for _, comment := range list {
		var replyItem model.ReplyCommentResp
		str, _ := json.Marshal(comment)
		json.Unmarshal(str,&replyItem)
		replyItem.ParentComment = comment.GetCommentById(comment.ParentCommentId)
		result = append(result, replyItem)
	}
	return result
}