package model

type Tag struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
}

type BlogTag struct {
	Id int64 `json:"id"`
	BlogId int64 `json:"blog_id"`
	TagId int64 `json:"tag_id"`
}

