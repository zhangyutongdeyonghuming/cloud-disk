package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// InitRouter 初始化路由并启动web
func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), log())
	err := r.Run()
	if err != nil {
		logrus.Error("web start error.", err)
		return
	}
	logrus.Info("web started!")
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

		// 日志格式
		logrus.WithFields(
			logrus.Fields{
				"latencyTime": latencyTime,
				"method":      method,
				"uri":         uri,
				"status":      status,
				"ip":          ip,
			}).Info()
	}
}
