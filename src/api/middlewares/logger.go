package middlewares

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"github.com/gin-gonic/gin"
)

func StructuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // start
		path := c.FullPath()
		raw:= c.Request.URL.RawQuery

		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now() // stop
		param.Latency = param.TimeStamp.Sub(start)
		
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()

		if raw != "" {
			param.Path = path + "?" + raw
		} else {
			param.Path = path
		}

		keys := map[logging.ExtraKey]interface{}{}

		keys[logging.ClientIP] = param.ClientIP
		keys[logging.Method] = param.Method
		keys[logging.Path] = param.Path
		keys[logging.StatusCode] = param.StatusCode
		keys[logging.ErrorMessage] = param.ErrorMessage
		keys[logging.BodySize] = param.BodySize
		keys[logging.Latency] = param.Latency
		keys[logging.RequestBody] = string(bodyBytes)
		// keys[logging.ResponseBody] = string(bodyBytes)


		logger.Info(logging.RequestResponse, logging.Api, "", keys)
	}
}