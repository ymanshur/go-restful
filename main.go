package main

import (
	"fmt"
	"go-restful/internal/factory"
	"go-restful/internal/http"
	"go-restful/pkg/constant"
	"go-restful/pkg/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init echo
	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	// Register validator
	e.Validator = util.NewValidator()

	f := factory.NewFactory()
	http.New(e.Group("/api"), f)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", constant.Env.Get("PORT", "8000"))))
}
