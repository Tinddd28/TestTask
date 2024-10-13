package handler

import (
	"github.com/Tinddd28/TestTask/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get all songs
// @Description Get all songs
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string false "Group name"
// @Param song query string false "Song name"
// @Param startDate query string false "Start date"
// @Param endDate query string false "End date"
// @Param offset query int false "Page number"
// @Success 200 {array} models.Song
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /songs [get]
func (h *Handler) GetAllSongs(c *gin.Context) {
	group := c.Query("group")
	song := c.Query("song")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	pageStr := c.Query("offset")

	page := 1
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, "invalid parameter")
			h.logger.Error("Error parsing page number", err)
			return
		}
	}

	songs, err := h.services.Song.GetAllSongs(group, song, startDate, endDate, page)
	if err != nil {
		h.logger.Error("Error getting songs", err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, songs)
	h.logger.Debug("Songs sent")
}

// @Summary Get song
// @Description Get song
// @Tags songs
// @Accept json
// @Produce json
// @Param id query int true "Song id"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} models.Verse
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /song [get]
func (h *Handler) GetSong(c *gin.Context) {
	idStr := c.Query("id")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	if idStr == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		h.logger.Error("Error getting id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, "invalid id")
		h.logger.Error("Error parsing id", err)
		return
	}

	page := 1
	pageSize := 10

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, "Invalid page number")
			h.logger.Error("Error parsing page number", err)
			return
		}
	}

	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 {
			c.JSON(http.StatusBadRequest, "Invalid page size")
			h.logger.Error("Error parsing page size", err)
			return
		}
	}

	verses, err := h.services.Song.GetSong(id, page, pageSize)
	if err != nil {
		h.logger.Error("Error getting verses", err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, verses)
	h.logger.Debug("Verses", verses)
}

// @Summary Delete song
// @Description Delete song
// @Tags songs
// @Accept json
// @Produce json
// @Param id query int true "Song id"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /song [delete]
func (h *Handler) DeleteSong(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		h.logger.Error("Error getting id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, "invalid id")
		h.logger.Error("Error parsing id", err)
		return
	}

	err = h.services.Song.DeleteSong(id)
	if err != nil {
		h.logger.Error("Error deleting song", err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.Status(http.StatusOK)
	h.logger.Debug("Song deleted")
}

// @Summary Update song
// @Description Update song
// @Tags songs
// @Accept json
// @Produce json
// @Param id query int true "Song id"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /song [put]
func (h *Handler) UpdateSong(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		h.logger.Error("Error getting id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, "invalid id")
		h.logger.Error("Error parsing id", err)
		return
	}

	var song models.Song
	err = c.ShouldBindJSON(&song)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request body")
		h.logger.Error("Error decoding request body", err)
		return
	}

	err = h.services.Song.UpdateSong(id, song)
	if err != nil {
		h.logger.Error("Error updating song", err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.Status(http.StatusOK)
	h.logger.Debug("Song updated")
}

// @Summary Create song
// @Description Create song
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.RequestSong true "Song info"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /song [post]
func (h *Handler) CreateSong(c *gin.Context) {
	var song models.RequestSong

	err := c.ShouldBindJSON(&song)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request body")
		h.logger.Error("Error decoding request body", err)
		return
	}

	id, err := h.services.CreateSong(song)
	if err != nil {
		h.logger.Error("Error creating song", err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
	h.logger.Debug("Song created")
}
