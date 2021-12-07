package access_control

import (
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/JieeiroSst/itjob/pkg/metric"
)

func Metrics(mService metric.UseCase) gin.HandlerFunc  {
	return func(c *gin.Context) {
		appMetric := metric.NewHTTP(c.Request.URL.Path, c.Request.Method)
		appMetric.Started()
		c.Next()
		    statusCode := c.Writer.Status()
			appMetric.Finished()
			appMetric.StatusCode = strconv.Itoa(statusCode)
			mService.SaveHTTP(appMetric)
	}
}