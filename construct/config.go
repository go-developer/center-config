// Package construct...
//
// Description : 初始化配置信息
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-23 10:52 下午
package construct

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig()  {
	viper.SetConfigType("yaml")
	v := viper.New()
	fmt.Println(v)
	v.SetConfigFile("")
	// 处理配置文件路径,读取优先级 env > local
	// configPath :=
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {   //注册文件监控的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
}