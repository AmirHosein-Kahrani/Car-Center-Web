package middlewares

import (
	"strconv"
	"time"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/metrics"
	"github.com/gin-gonic/gin"
)

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		if path == "" {
			path = "/undifined"
		}
		method := c.Request.Method
		c.Next()
		status := c.Writer.Status()
		// * another way to create a timer
		// timer := prometheus.NewTimer(metrics.HttpDuration.WithLabelValues(path, method, strconv.Itoa(status)))
		// timer.ObserveDuration()
		metrics.HttpDuration.WithLabelValues(path, method, strconv.Itoa(status)).
			Observe(float64(time.Since(start) / time.Millisecond))

	}
}
