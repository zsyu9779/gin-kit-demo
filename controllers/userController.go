package controllers

import (
	"gin-kit-demo/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var userService service.UserService
func Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	_,isLogin:=userService.Login(session,username,password)
	if isLogin {
		c.JSON(200,0)
	}else {
		c.JSON(200,1)
	}
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	c.JSON(200,1)
}
func GetUser(c *gin.Context) {
	session := sessions.Default(c)
	user :=session.Get("user")
	c.JSON(200,user)
}
func IsLogin(c *gin.Context)  {
	session := sessions.Default(c)
	user :=session.Get("user")
	if user != nil {
		c.JSON(200,1)
	}else {
		c.JSON(200,0)
	}

}

func GetUsers(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func AddUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
