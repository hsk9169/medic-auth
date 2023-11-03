package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medic-basic/auth/pkg/domain/model"
	"github.com/medic-basic/auth/pkg/handler/common"
)

const (
	SignInPath = "/signin"
)

type SignInRequest struct {
	Phone string `form:"phone"`
	Role  string `form:"role"`
}

type SignInResponse struct {
	AuthData    model.TokenDetailsParam `json:"authData"`
	AccountData any                     `json:"accountData"`
}

type SignInHandler struct{}

func (h SignInHandler) GetPathHttpMethod() (string, string, bool) {
	return SignInPath, http.MethodPost, false
}

// Handle godoc
// @Summary sign in
// @Description create Token by phone number
// @Accept json
// @Produce json
// @param SignInRequest query SignInRequest true "Token object"
// @Router /auth/signin [post]
func (h SignInHandler) Handle(c *gin.Context) {
	var req SignInRequest
	if err := common.BindReqParamsAndSetInCtx(c, &req); err != nil {
		common.SetErrorResponse(c, err)
		return
	}

	auth, err := common.GetTokenAggregate(c)
	if err != nil {
		common.SetErrorResponse(c, err)
		return
	}

	userVerifyParam := model.UserVerifyParam{
		Phone: req.Phone,
		Role:  req.Role,
	}

	accountData, err := auth.VerifyUser(c, userVerifyParam)
	if err != nil {
		common.SetUnauthorizedResponse(c)
		return
	}

	tokens, err := auth.CreateToken(c, userVerifyParam)
	if err != nil {
		common.SetErrorResponse(c, err)
		return
	}

	res := SignInResponse{
		AuthData:    *tokens,
		AccountData: accountData,
	}

	c.JSON(http.StatusOK, res)
}
