package services

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetFavorites(c echo.Context) error {
	slog.Info("Getting favorites")
	favs, err := selectAllFavorites()
	if err != nil {
		slog.Error("Error getting favorites", "selectAllFavorites", err)
		return c.JSON(500, "Error getting favorites")
	}
	return c.JSON(200, favs)
}

func InsertFavorite(c echo.Context) error {
	slog.Info("Inserting favorite")

	fav := new(Favorite)

	if err := c.Bind(fav); err != nil {
		slog.Error("Error binding request body", "error", err)
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	count, countErr := countFavorites()

	if countErr != nil {
		slog.Error("Error counting favorites", "error", countErr)
		return c.JSON(http.StatusInternalServerError, "Error counting favorites")
	}

	if count >= 5 {
		return c.JSON(http.StatusForbidden, "Cannot insert more than 5 favorites")
	}

	err := insertFavorite(fav)
	if err != nil {
		slog.Error("Error inserting favorite", "insertFavorite", err)
		return c.JSON(http.StatusInternalServerError, "Error inserting favorite")
	}

	slog.Info("Favorite inserted successfully", "favorite", fav)
	return c.JSON(http.StatusCreated, fav)
}

func DeleteFavorite(c echo.Context) error {
	slog.Info("Deleting favorite")

	malIDStr := c.Param("mal_id")
	malID, err := strconv.Atoi(malIDStr)

	if err != nil {
		slog.Error("Invalid mal_id format", "mal_id", malIDStr)
		return c.JSON(http.StatusBadRequest, "Invalid mal_id format")
	}

	exists, err := malIDExists(malID)

	if err != nil {
		slog.Error("Error checking mal_id existence", "error", err)
		return c.JSON(http.StatusInternalServerError, "Error checking mal_id")
	}

	if !exists {
		slog.Warn("Favorite with mal_id not found", "mal_id", malID)
		return c.JSON(http.StatusNotFound, "Favorite not found on the list")
	}

	err = deleteFavoriteByMalID(malID)
	if err != nil {
		slog.Error("Error deleting favorite", "deleteFavoriteByMalID", err)
		return c.JSON(http.StatusInternalServerError, "Error deleting favorite")
	}

	slog.Info("Favorite deleted successfully", "mal_id", malID)
	return c.JSON(http.StatusOK, "Favorite deleted successfully")
}
