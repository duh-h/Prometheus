package routes

import (
	"net/http"
	"time"

	"github.com/duh-h/Prometheus/api/database"
	"github.com/duh-h/Prometheus/api/models"
	"github.com/gin-gonic/gin"
)

func EditURL(c *gin.Context) {
	shortID := c.Param("shortID")

	var boby models.Request

	if err := c.ShouldBind(&boby); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot parse JSON",
		})
		return
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()

	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "short id not exists",
		})
		return
	}

	err = r.Set(database.Ctx, shortID, boby.URL, boby.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to update the shortend link",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "the content has been update",
	})

}
