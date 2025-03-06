package routes

import (
	"net/http"

	"github.com/duh-h/Prometheus/api/database"
	"github.com/gin-gonic/gin"
)

func DeleteURL(c *gin.Context) {
	shortID := c.Param("shortID")

	//r := database.CreateClient(0)
	//defer r.Close()

	err := database.Client.Del(database.Ctx, shortID).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete shortend link",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Shortend URL delete successfull ",
	})
}
