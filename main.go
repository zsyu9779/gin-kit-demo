package main

import (
	"fmt"
	"gin-kit-demo/conf"
	"gin-kit-demo/controllers"
	"gin-kit-demo/g_rediscache"
	"gin-kit-demo/model"
	"gin-kit-demo/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	//pflag.Int("flagname", 1234, "help message for flagname")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	conf.InitConfig()
	router.Include(controllers.Routers)
	// 初始化路由
	r := router.Init()
	model.Init()
	defer model.Db.Close()
	// 初始化gorm
	//model.Init()
	// 初始化redis
	g_rediscache.Init()
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
