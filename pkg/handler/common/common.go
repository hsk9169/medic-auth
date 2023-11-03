package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/medic-basic/auth/pkg/domain/model"
	"github.com/pkg/errors"
)

const (
	APIReqParamsKeyName = "apiParams"

	AuthKeyName  = "AuthKey"
	QueryKeyName = "QueryKey"
)

func BindReqParamsAndSetInCtx(c *gin.Context, reqParams any) error {
	if err := c.ShouldBind(reqParams); err != nil {
		return err
	}
	b, _ := jsoniter.Marshal(reqParams)
	c.Set(APIReqParamsKeyName, string(b))
	return nil
}

func BindReqUriParamsAndSetInCtx(c *gin.Context, reqParams any) error {
	if err := c.ShouldBindUri(reqParams); err != nil {
		return err
	}
	return nil
}

func SetErrorResponse(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func SetUnauthorizedResponse(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
}

func GetTokenAggregate(c *gin.Context) (model.AuthAggregate, error) {
	val, exists := c.Get(AuthKeyName)
	if !exists {
		return nil, errors.New("auth aggregate not exists")
	}
	return val.(model.AuthAggregate), nil
}
