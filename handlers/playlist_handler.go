package handlers

import (
	"music-streaming/database"
	"music-streaming/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreatePlaylist - создание нового плейлиста
func CreatePlaylist(c echo.Context) error {
	playlist := new(models.Playlist)
	if err := c.Bind(playlist); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Неверные данные"})
	}

	_, err := database.DB.Exec("INSERT INTO playlists (name, user_id) VALUES ($1, $2)", playlist.Name, playlist.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Ошибка создания плейлиста"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Плейлист создан"})
}

// GetPlaylists - получение всех плейлистов пользователя
func GetPlaylists(c echo.Context) error {
	userID := c.QueryParam("user_id")
	rows, err := database.DB.Query("SELECT id, name FROM playlists WHERE user_id = $1", userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Ошибка получения плейлистов"})
	}
	defer rows.Close()

	var playlists []models.Playlist
	for rows.Next() {
		var playlist models.Playlist
		rows.Scan(&playlist.ID, &playlist.Name)
		playlists = append(playlists, playlist)
	}

	return c.JSON(http.StatusOK, playlists)
}

// DeletePlaylist - удаление плейлиста
func DeletePlaylist(c echo.Context) error {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM playlists WHERE id = $1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Ошибка удаления плейлиста"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Плейлист удален"})
}
