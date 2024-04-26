package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {

	Log.SetFormatter(&logrus.TextFormatter{})
	Log.SetReportCaller(true)
	Log.Level = logrus.DebugLevel
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		userAgent := c.Request.UserAgent()

		Log.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency":     latency,
			"client_ip":   clientIP,
			"method":      method,
			"path":        path,
			"user_agent":  userAgent,
		}).Debugln("handled request")
	}
}

func TraceId(id, message string) {

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Log.WithField("id", id).Debug(message)
}
