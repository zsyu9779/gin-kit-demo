package model

//easyjson:json
type CommentResp struct {
	Id            int64         `json:"id"`
	User          string        `json:"user"`
	Avatar        string        `json:"avatar"`
	Content       string        `json:"content"`
	ReplyComment  []CommentResp `json:"reply_comment"`
	ParentComment *CommentResp  `json:"parent_comment"`
	Ctime         int64         `json:"ctime"`
}

//easyjson:json
type Comment struct {
	Id       int64   `json:"id"`
	User     string  `json:"user"`
	Avatar   string  `json:"avatar"`
	Content  string  `json:"content"`
	ParentId int64   `json:"parent_id"`
	ReplyIds []int64 `json:"reply_ids"`
	Ctime    int64   `json:"ctime"`
}
