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
	RefreshTokenPath = "/refresh"
)

type RefreshTokenResponse struct {
	AuthData model.TokenDetailsParam `json:"authData"`
}

type RefreshTokenHandler struct{}

func (h RefreshTokenHandler) GetPathHttpMethod() (string, string, bool) {
	return RefreshTokenPath, http.MethodGet, true
}

// Handle godoc
// @Summary refresh Tokens
// @Description refresh Tokens
// @Accept json
// @Produce json
// @Router /auth/refresh [get]
func (h RefreshTokenHandler) Handle(c *gin.Context) {
	auth, err := common.GetTokenAggregate(c)
	if err != nil {
		common.SetErrorResponse(c, err)
		return
	}

	authToken, isValid := c.Get("auth_token")
	if !isValid {
		common.SetUnauthorizedResponse(c)
		return
	}

	claims, err := util.GetJWTClaims(authToken)
	if err != nil {
		fmt.Println(err)
		common.SetUnauthorizedResponse(c)
		return
	}

	if err := auth.CheckTokenValid(c, model.TokenClaimsParam{
		Phone: claims.Phone,
		Role:  claims.Role,
		Uuid:  claims.Uuid,
	}); err != nil {
		common.SetUnauthorizedResponse(c)
		return
	}

	userVerifyParam := model.UserVerifyParam{
		Phone: claims.Phone,
		Role:  claims.Role,
	}

	_, err = auth.VerifyUser(c, userVerifyParam)
	if err != nil {
		common.SetUnauthorizedResponse(c)
		return
	}

	tokens, err := auth.CreateToken(c, userVerifyParam)
	if err != nil {
		common.SetErrorResponse(c, err)
		return
	}

	res := RefreshTokenResponse{
		AuthData: *tokens,
	}

	c.JSON(http.StatusOK, res)
}
