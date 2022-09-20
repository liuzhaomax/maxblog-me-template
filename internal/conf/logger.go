package conf

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"maxblog-me-template/internal/core"
	"maxblog-me-template/internal/utils"
	"os"
	"time"
)

func init() {
	logrus.SetLevel(logrus.InfoLevel) // Trace << Debug << Info << Warning << Error << Fatal << Panic
	InitializeLogging("golog.txt")    // TODO 根据时间创建不同的日志文件，减小IO开支
}

func InitializeLogging(logFile string) {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"失败方法": utils.GetFuncName(),
		}).Panic(core.FormatError(902, err).Error())
	}
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func LoggerToFile() gin.HandlerFunc {
	fileName := "golog.txt"
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"失败方法": utils.GetFuncName(),
		}).Panic(core.FormatError(902, err).Error())
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %8s | %s ",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
