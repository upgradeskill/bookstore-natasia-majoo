package server

import (
	"bookstore/api"
	"bookstore/util"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes() {
	e := echo.New()

	v1 := e.Group("/v1")

	// for public (unauthorized user)
	{
		api.Authorization(v1)
		api.BookAPI(v1)
	}

	v1.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(util.GlobalConfig.JWT_SECRET),
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	}))

	// for authorized user
	{
		api.AdminBookAPI(v1)
	}

	if err := e.Start(":8000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
