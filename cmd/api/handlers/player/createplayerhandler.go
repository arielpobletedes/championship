package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/internal/domain"
)

func (h Handler) CreatePlayerHandler(c *gin.Context) {
	var player domain.Player
	if err := c.BindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertResult, err := h.PlayerService.InsertPlayer(player)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops, something went wrong!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"player_id": insertResult,
	})
}
