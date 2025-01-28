package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JWTMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("your_secret_key"),
})
