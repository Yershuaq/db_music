package routes

import (
	"music-streaming/handlers"
	"music-streaming/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Аутентификация
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	// Плейлисты (JWT защита)
	playlistGroup := e.Group("/playlists")
	playlistGroup.Use(middlewares.JWTMiddleware)
	playlistGroup.POST("", handlers.CreatePlaylist)
	playlistGroup.GET("", handlers.GetPlaylists)
	playlistGroup.DELETE("/:id", handlers.DeletePlaylist)

	// Поиск песен
	e.GET("/songs", handlers.SearchSongs)
}
