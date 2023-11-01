package main

import (
	"goVanila/internal/database"

	"goVanila/internal/delivery/logger"
	"goVanila/internal/delivery/rest"

	"goVanila/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	logger.Init()
	e := echo.New()
	rest.LoadMiddlewares(e)

	db := database.GetDb()

	container := usecase.NewContainer(db)
	h := rest.NewHandler(container.PgUtil, container.AuthUsecase, container.Middleware, db)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start((":4000")))

}
