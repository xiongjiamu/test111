package httpserver

import (
	"dkdns/dkFramework/configs"
	"dkdns/dkFramework/logger"
	"dkdns/httpServer/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartGinServer(appConfigs *configs.Config) {
	router := gin.Default()

	// 设置静态文件目录
	router.StaticFS("/assets", http.Dir(appConfigs.HTTPS.ServerPath+"/views/assets")) // 将 assert 文件夹及其子文件夹设置为静态资源文件夹
	// 设置路由
	router.GET("/", controllers.DefaultHandler)
	router.GET("/cachelist", controllers.CacheListHandler)
	router.POST("/addspecial", controllers.AddSpecialHandler)
	router.DELETE("/delspecial", controllers.DelSpecialHandler)
	router.DELETE("/delcache", controllers.DelCacheHandler)

	logger.Info("http://localhost:8080")
	// 启动服务器
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
