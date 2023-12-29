package main

import (
	"dkdns/dkFramework/configs"
	"dkdns/dkFramework/logger"
	dkdnserver "dkdns/dnsServer"
)

func main() {
	// 加载配置文件
	config, err := configs.LoadConfig("etc/dkdns.yaml")
	if err != nil {
		panic(err)
	}

	// 初始化日志记录器
	logger.Init(config)
	logger.Println("SysConfig：  ", config)

	// 启动DNS Server
	go dkdnserver.RunDNSServer(config)

	// 主协程等待，否则程序会立即退出
	select {}
}
