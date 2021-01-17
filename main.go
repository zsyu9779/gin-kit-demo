package main

import (
	"gin-kit-demo/conf"
	"gin-kit-demo/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	//pflag.Int("flagname", 1234, "help message for flagname")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	conf.InitConfig()
	router.Include()
}
