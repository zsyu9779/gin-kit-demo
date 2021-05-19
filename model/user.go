package model

type User struct {
	Id         int64  `json:"id"`
	Avatar     string `json:"avatar"`
	CreateTime string `json:"create_time"`
	Email      string `json:"email"`
	NickName   string `json:"nickname"`
	Password   string `json:"password"`
	Type       int    `json:"type"`
	UpdateTime string `json:"update_time"`
	UserName   string `json:"username"`
}
