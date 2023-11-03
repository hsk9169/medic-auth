package log

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/medic-basic/auth/pkg/handler/common"
	"github.com/sirupsen/logrus"
)

const (
	TypeInfo   = "info"
	TypeError  = "error"
	TypeResult = "result"

	TraceIdKeyName = "trace_id"

	logTypeKeyName      = "log_type"
	apiRequestParameter = "api_request_parameter"
)

func Infof(ctx context.Context, format string, args ...any) {
	logrus.WithField(logTypeKeyName, TypeInfo).WithField(TraceIdKeyName, ctx.Value(TraceIdKeyName)).Infof(format, args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	logrus.WithField(logTypeKeyName, TypeError).WithField(TraceIdKeyName, ctx.Value(TraceIdKeyName)).Errorf(format, args...)
}

func Resultf(ctx context.Context, format string, args ...any) {
	logrus.WithField(logTypeKeyName, TypeResult).WithField(TraceIdKeyName, ctx.Value(TraceIdKeyName)).Infof(format, args...)
}

func GetResultLogger(c *gin.Context) *logrus.Entry {
	return logrus.WithField(logTypeKeyName, TypeResult).
		WithField(TraceIdKeyName, c.Value(TraceIdKeyName)).
		WithField(apiRequestParameter, c.GetString(common.APIReqParamsKeyName))
}
