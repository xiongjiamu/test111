package controllers

import (
	"dkdns/dkFramework/databases"
	"dkdns/dkFramework/logger"
	"github.com/gin-gonic/gin"
	"html/template"
	"net"
	"net/http"
	"regexp"
)

func DefaultHandler(c *gin.Context) {

	specicalMap, err := databases.Boltdb.ReadBucket("special_list")

	// 渲染模板并将数据传递给模板
	tmpl, err := template.ParseFiles("httpserver/views/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing template")
		return
	}

	// 如果需要在模板中使用读取到的数据，可以将其放入一个结构体中，然后传递给模板
	type ViewData struct {
		YourData map[string]string // 请根据你的数据结构定义这个字段
	}

	viewData := ViewData{
		YourData: specicalMap, // 假设数据是字符串类型，将字节数组转换为字符串
	}

	err = tmpl.Execute(c.Writer, viewData)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error executing template")
		return
	}
}

func IsValidDomain(domain string) bool {
	// 校验 domain 是否为域名格式
	domainRegex := `^([a-zA-Z0-9_-]+\.)+[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(domainRegex, domain)
	return match
}

func IsValidIP(ip string) bool {
	// 校验 IP 是否为合法的 IPv4 或 IPv6 地址格式
	return net.ParseIP(ip) != nil
}

func DelSpecialHandler(c *gin.Context) {

	key := c.Query("key")

	err := databases.Boltdb.Delete("special_list", key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
	})

}

func AddSpecialHandler(c *gin.Context) {

	type SpecialData struct {
		Domain string `json:"domain"`
		IP     string `json:"ip"`
	}
	var specialData SpecialData

	if err := c.BindJSON(&specialData); err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON format")
		return
	}

	domain := specialData.Domain
	ip := specialData.IP

	logger.Println(domain, "dfsdfsdfs")
	logger.Info(ip)

	if !IsValidDomain(domain) {
		c.String(http.StatusOK, "Invalid domain format")
		return
	}

	if !IsValidIP(ip) {
		c.String(http.StatusBadRequest, "Invalid IP format")
		return
	}

	err := databases.Boltdb.Write("special_list", domain, ip)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
	})
}

// 添加更多的处理函数...
