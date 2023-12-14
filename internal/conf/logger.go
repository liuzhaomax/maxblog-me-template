package conf

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"maxblog-me-template/internal/core"
	"os"
	"time"
)

func InitLogger() {
	logFile := "golog.txt" // TODO 根据时间创建不同的日志文件，减小IO开支
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"失败方法": core.GetFuncName(),
		}).Panic(core.FormatError(902, err).Error())
	}
	logrus.SetLevel(logrus.InfoLevel) // Trace << Debug << Info << Warning << Error << Fatal << Panic
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: cfg.Logger.Color})
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
}

type LoggerFormat struct {
	StatusCode int           `json:"code"`
	Took       time.Duration `json:"took"`
	ClientIP   string        `json:"client_ip"`
	Method     string        `json:"method"`
	URI        string        `json:"uri"`
}

func LoggerToFile() gin.HandlerFunc {
	fileName := "golog.txt"
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"失败方法": core.GetFuncName(),
		}).Panic(core.FormatError(902, err).Error())
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		took := endTime.Sub(startTime)
		method := c.Request.Method
		uri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		//logger.Infof("| %3d | %13v | %15s | %8s | %s ",
		//    statusCode,
		//    took,
		//    clientIP,
		//    method,
		//    uri,
		//)
		format := &LoggerFormat{
			StatusCode: statusCode,
			Took:       took,
			ClientIP:   clientIP,
			Method:     method,
			URI:        uri,
		}
		formatBytes, _ := json.Marshal(format)
		logger.Infof(string(formatBytes))
	}
}
