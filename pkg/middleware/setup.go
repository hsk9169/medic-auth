package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/medic-basic/auth/pkg/domain"
	"github.com/medic-basic/auth/pkg/log"
)

// Setup 호출시 제어 가능한 첫번째 함수 구간으로 새로운 호출에 대한 기본 셋업을 진행한다.
func Setup() gin.HandlerFunc {
	return func(c *gin.Context) {
		if IsIgnoredMiddlewareReqURI(c.Request.RequestURI) {
			return
		}

		// Note: 가장 먼저 호출되는 미들웨어에서 API 호출 시작 시간을 저장한다.
		c.Set(apiStartTimeKeyName, time.Now())

		tid := uuid.New().String()
		c.Set(log.TraceIdKeyName, tid)

		domain.SetupDomain(c)
		c.Next()
	}
}
