package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	// 实例化
	logger := logrus.New()
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)
	rotateFileHook, _ := NewRotateFileHook(*Generate(logrus.DebugLevel))
	rotateFileHook2, _ := NewRotateFileHook(*Generate(logrus.InfoLevel))
	rotateFileHook3, _ := NewRotateFileHook(*Generate(logrus.WarnLevel))
	rotateFileHook4, _ := NewRotateFileHook(*Generate(logrus.ErrorLevel))
	logrus.AddHook(rotateFileHook)
	logrus.AddHook(rotateFileHook2)
	logrus.AddHook(rotateFileHook3)
	logrus.AddHook(rotateFileHook4)
	return func(c *gin.Context) {
		//10.1.1.1 - - [22/Aug/2014:16:48:14 +0800] "POST /ajax/MbpRequest.do HTTP/1.1" 200 367 "-" "Dalvik/1.6.0 (Linux; U; Android 4.1.1; ARMM7K Build/JRO03H)" "119.189.56.175" 127.0.0.1:8090 0.022 0.022

		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		//请求方式
		upgrade := c.Request.Header.Get("Upgrade")
		UserAgent := c.Request.Header.Get("User-Agent")
		length:=c.Request.Header.Get("Content-Length")
		referer := c.Request.Header.Get("Referer")
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			auth = "-"
		}
		requestDate := startTime.Format("02/Jan/2006:15:04:00 -0700")
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("%s %s - [%s] \"%s %s %s\" %d %s %f \"%s\" \"%s\"",clientIP,auth,requestDate,reqMethod,reqUrl,upgrade,statusCode,length,latencyTime.Seconds(),referer,UserAgent)

	}
}

func Generate(level logrus.Level) *RotateFileConfig {
	rotateFileConfig := RotateFileConfig{
		Filename:   fmt.Sprintf("../material_feed/log/%s-%s.log", level.String(), time.Now().Format(time.RFC3339)),
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     7,
		Level:      level,
		Formatter:  &logrus.JSONFormatter{},
	}
	writer, _ := rotatelogs.New(
		rotateFileConfig.Filename+".%Y-%m-%d-%H-%M",
		// 生成软链，指向最新日志文
		rotatelogs.WithLinkName(rotateFileConfig.Filename),
		// 文件最大保存时间
		rotatelogs.WithMaxAge(time.Duration(2)*time.Hour),
		// 日志切割时间间隔
		rotatelogs.WithRotationTime(time.Duration(5)*time.Second),
	)
	logrus.SetOutput(writer)
	return &rotateFileConfig
}
