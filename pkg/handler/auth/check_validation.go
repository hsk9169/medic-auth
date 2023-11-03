package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medic-basic/auth/pkg/domain/model"
	"github.com/medic-basic/auth/pkg/handler/common"
	"github.com/medic-basic/auth/pkg/util"
)

const (
	CheckValidationPath = "/valid"
)

type CheckValidationHandler struct{}

func (h CheckValidationHandler) GetPathHttpMethod() (string, string, bool) {
	return CheckValidationPath, http.MethodGet, true
}

// Handle godoc
// @Summary check Token Validation
// @Description check Token Validation
// @Accept json
// @Produce json
// @Router /auth/valid [get]
func (h CheckValidationHandler) Handle(c *gin.Context) {
	authToken, isValid := c.Get("auth_token")
	if !isValid {
		fmt.Println("no auth token in context")
		common.SetUnauthorizedResponse(c)
		return
	}
	claims, err := util.GetJWTClaims(authToken)
	if err != nil {
		fmt.Println(err)
		common.SetUnauthorizedResponse(c)
		return
	}

	auth, err := common.GetTokenAggregate(c)
	if err != nil {
		fmt.Println("failed to get token aggregate")
		common.SetErrorResponse(c, err)
		return
	}

	if err := auth.CheckTokenValid(c, model.TokenClaimsParam{
		Phone: claims.Phone,
		Role:  claims.Role,
		Uuid:  claims.Uuid,
	}); err != nil {
		fmt.Println("no token cache hit")
		common.SetUnauthorizedResponse(c)
		return
	}

	_, err = auth.VerifyUser(c, model.UserVerifyParam{
		Phone: claims.Phone,
		Role:  claims.Role,
	})
	if err != nil {
		fmt.Println("no user verified")
		common.SetUnauthorizedResponse(c)
		return
	}

	c.JSON(http.StatusOK, "success")
}
