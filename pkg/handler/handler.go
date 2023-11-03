package handler

import (
	"github.com/gin-gonic/gin"
	auth "github.com/medic-basic/auth/pkg/handler/auth"
)

type Handler interface {
	Handle(*gin.Context)
	GetPathHttpMethod() (string, string, bool)
}

type HandlerList []Handler

func GetHandlerList() HandlerList {
	return []Handler{
		HealthCheckHandler{},

		auth.CheckValidationHandler{},
		auth.RefreshTokenHandler{},
		auth.SignInHandler{},
		auth.SignOutHandler{},
	}
}
