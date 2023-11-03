package middleware

import (
	"bytes"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medic-basic/auth/pkg/log"
	"github.com/sirupsen/logrus"
)

const (
	logTypeResult       = "result"
	bodyKey             = "body"
	apiStartTimeKeyName = "api_start_time"
)

type BodyLowWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLowWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := &bytes.Buffer{}
		c.Set(bodyKey, body)
		c.Writer = &BodyLowWriter{body: body, ResponseWriter: c.Writer}
		c.Next()

		logResult(c, time.Since(c.GetTime(apiStartTimeKeyName)), body)
	}
}

func logResult(c *gin.Context, time time.Duration, body *bytes.Buffer) {

	resultLogger := log.GetResultLogger(c).WithFields(logrus.Fields{
		"method":  c.Request.Method,
		"status":  c.Writer.Status(),
		"api":     c.FullPath(),
		"latency": time,
	})

	result := strings.Replace(body.String(), "\n", "", -1)

	if errMsg := c.Errors.String(); len(errMsg) > 0 {
		resultLogger.WithFields(logrus.Fields{"errMsg": errMsg, "result": result}).Error(result)
	} else {
		resultLogger.WithFields(logrus.Fields{"result": result}).Info(result)
	}

}

func IsIgnoredMiddlewareReqURI(reqURI string) bool {
	for _, ignoredPath := range []string{"/health/", "/swagger/"} {
		if strings.HasPrefix(reqURI, ignoredPath) {
			return true
		}
	}

	return false
}
