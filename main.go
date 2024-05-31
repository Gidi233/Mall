package main

import (
	"Mall/core"
	"Mall/global"
)

func main() {

	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库

}
