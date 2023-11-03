package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medic-basic/auth/pkg/handler"
	"github.com/medic-basic/auth/pkg/middleware"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

const (
	swaggerPath     = "/swagger/*any"
	shutdownTimeout = 5
)

type Server struct {
	*http.Server
}

func NewServer(port int, version string) *Server {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())
	router.Use(middleware.Setup())
	router.Use(middleware.Log())

	for _, cmd := range handler.GetHandlerList() {
		path, method, mwRequired := cmd.GetPathHttpMethod()
		if mwRequired {
			router.Handle(method, path, middleware.TokenAuthMiddleware(), cmd.Handle)
		} else {
			router.Handle(method, path, cmd.Handle)
		}
	}

	// swagger
	docs.SwaggerInfo.Title = "auth"
	docs.SwaggerInfo.Description = "auth " + version
	router.GET(swaggerPath, ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &Server{
		Server: &http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: router,
		},
	}
}

func (s *Server) Run() error {
	return s.ListenAndServe()
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}
}
