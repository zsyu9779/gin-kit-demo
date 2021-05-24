package service

import (
	"gin-kit-demo/model"
	"github.com/gin-contrib/sessions"
)

type UserService struct {
}

func (u *UserService) Login(session sessions.Session,username, password string) (model.User,bool) {
	user := model.User{UserName: username}
	user.GetUserByName()
	if user.Password != password {
		return user,false
	}
	session.Set("user",user)
	session.Save()
	return user,true
}

func (u *UserService) FindUserById(userId int64) model.User{
	user := model.User{Id: userId}
	user.GetUserById()
	return user
}