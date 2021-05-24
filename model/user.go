package model

import "github.com/sirupsen/logrus"

type User struct {
	Id         int64  `json:"id"`
	Avatar     string `json:"avatar"`
	CreateTime string `json:"createTime"`
	Email      string `json:"email"`
	NickName   string `json:"nickname"`
	Password   string `json:"password"`
	Type       int    `json:"type"`
	UpdateTime string `json:"updateTime"`
	UserName   string `json:"username"`
}


func (u *User)GetUserByName() *User {
	err := Db.Where("user_name = ?",u.UserName).First(u).Error
	if err !=nil {
		return nil
	}
	return u
}
func (u *User)GetUserById()  {
	err := Db.First(u,u.Id).Error
	if err !=nil {
		logrus.Error(err)
		panic(err)
	}
}
func (t *User)DeleteUser() bool {
	num := Db.Delete(&t,t.Id).RowsAffected
	if num>0 {
		return true
	}else {
		return false
	}
}

func (t *Type)UpdateUser() bool {
	num := Db.Save(&t).RowsAffected
	if num>0 {
		return true
	}else {
		return false
	}
}
func (t *Type)AddUser() bool {
	num := Db.Create(&t).RowsAffected
	if num>0 {
		return true
	}else {
		return false
	}
}