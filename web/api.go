// Package web 提供webapi接口调用
package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kunpeng/config"
	"kunpeng/plugin"
)

// StartServer 启动web服务接口
func StartServer(bindAddr string) {
	router := gin.Default()
	router.GET("/api/pluginList", func(c *gin.Context) {
		c.JSON(200, plugin.GetPlugins())
	})
	router.POST("/api/check", func(c *gin.Context) {
		var json plugin.Task
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := plugin.Scan(json)
		c.JSON(200, result)
	})
	router.POST("/api/config", func(c *gin.Context) {
		buf := make([]byte, 2048)
		n, _ := c.Request.Body.Read(buf)
		config.Set(string(buf[0:n]))
		c.JSON(200, map[string]bool{"success": true})
	})

	router.Run(bindAddr)
}
