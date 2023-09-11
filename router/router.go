package router

import (
	"cloud-disk/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// InitRouter 初始化路由并启动web
func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), log())
	port := config.Cfg.Server.Port
	if err := r.Run(":" + port); err != nil {
		logrus.Error("web start error.", err)
		return
	}
}

// log gin请求日志记录
func log() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		method := c.Request.Method

		// 请求路由
		uri := c.Request.RequestURI

		// 状态码
		status := c.Writer.Status()

		// 请求IP
		ip := c.ClientIP()
		data := map[string]interface{}{
			"latencyTime": latencyTime,
			"method":      method,
			"uri":         uri,
			"status":      status,
			"ip":          ip,
		}
		// 日志格式
		logrus.WithField("data", data).Info()
	}
}
