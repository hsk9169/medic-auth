package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medic-basic/auth/pkg/handler/common"
	"github.com/medic-basic/auth/pkg/util"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		tokenStr := util.ExtractToken(bearerToken)
		token, err := util.VerifyAuthToken(tokenStr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("middleware auth error")
			common.SetUnauthorizedResponse(c)
			c.Abort()
			return
		}
		c.Set("auth_token", token)
		c.Next()
	}
}
