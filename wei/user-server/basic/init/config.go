package init

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/yuhang-jieke/week2/wei/user-server/basic/config"
)

func ConfigInit() {
	viper.SetConfigFile("C:\\Users\\ZhuanZ\\Desktop\\week2\\week2\\wei\\user-server\\dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config.GlobalConf)
	if err != nil {
		return
	}
	fmt.Println("配置文件读取成功")
}
