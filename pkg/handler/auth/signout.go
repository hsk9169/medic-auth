package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medic-basic/auth/pkg/handler/common"
)

const (
	SignOutPath = "/signout"
)

type SignOutRequest struct {
	Phone string `form:"phone"`
}

type SignOutResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignOutHandler struct{}

func (h SignOutHandler) GetPathHttpMethod() (string, string, bool) {
	return SignOutPath, http.MethodPost, true
}

// Handle godoc
// @Summary sign in
// @Description create Token by phone number
// @Accept json
// @Produce json
// @param SignInRequest query SignInRequest true "Token object"
// @Router /auth/signin [post]
func (h SignOutHandler) Handle(c *gin.Context) {
	var req SignOutHandler
	if err := common.BindReqParamsAndSetInCtx(c, &req); err != nil {
		common.SetErrorResponse(c, err)
		return
	}

	// make query

	var resp SignOutResponse = SignOutResponse{}

	c.JSON(http.StatusOK, resp)
}
