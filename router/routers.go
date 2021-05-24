package router

import (
	"encoding/gob"
	"fmt"
	"gin-kit-demo/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"time"
)

type Option func(*gin.Engine)

var options []Option

func Include(option ...Option) {
	options = append(options, option...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// todo:调研gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	gob.Register(model.User{})
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %s %d %s \"%s\" %s\"\n",
			param.Request.RemoteAddr,             // [::1]:63495
			param.ClientIP,                       // [::1]
			param.TimeStamp.Format(time.RFC3339), //"2020-06-05T22:15:30+08:00
			param.Method,                         //GET
			param.Path,                           //  /v1/user/6
			param.Request.Proto,                  // =HTTP/1.1
			param.StatusCode,                     // 200
			param.Latency,                        // 5.24106793s
			param.Request.UserAgent(),            // PostmanRuntime/7.6.0
			param.ErrorMessage,                   // ""
		)
	}))
	r.Use(Cors())
	r.Use(gin.Recovery())
	for _, opt := range options {
		opt(r)
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
