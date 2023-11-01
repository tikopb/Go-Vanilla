package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadRoutes(e *echo.Echo, handler *handler) {
	//userAuth
	userGroup := e.Group("/user")
	userGroup.POST("/register", handler.RegisterUser)
	userGroup.POST("/login", handler.Login)
	userGroup.GET("/refresh", handler.RefreshSession)
	userGroup.GET("/logout", handler.LogOut)

	//cors test
	e.GET("/cors", handler.CorsTest)
}

func LoadMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
	}))
}
