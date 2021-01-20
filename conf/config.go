package conf

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
)

func InitConfig()error {
	/*
		func Caller(skip int) (pc uintptr, file string, line int, ok bool)
		skip是要提升的堆栈帧数，0-当前函数，1-上一层函数，....
		返回值：pc是uintptr这个返回的是函数指针，file是函数所在文件名目录，line所在行号，ok 是否可以获取到信息
	*/
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	logrus.Info(fmt.Sprintf("Filename : %q, Dir : %q\n", filename, path.Dir(filename)))
	box := packr.New("config", path.Dir(filename))
	dedaultConfig, _ := box.Find("default.yml")

	v := viper.New()
	v.SetConfigType("yml")
	err := v.ReadConfig(bytes.NewReader(dedaultConfig))
	if err != nil {
		panic(fmt.Sprintf("load conf has err:%s", err))
	}
	//写入默认配置
	configs := v.AllSettings()
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	env := os.Getenv("GO_ENV")
	logrus.Info("env is", env)
	if env != "" {
		envConfig, _ := box.Find(env + ".yml")
		viper.SetConfigType("yml")
		err = viper.ReadConfig(bytes.NewReader(envConfig))
		if err != nil {
			return nil
		}
	}
	return nil
}
