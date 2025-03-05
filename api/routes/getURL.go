package routes

import (
	"net/http"

	"github.com/duh-h/Prometheus/api/database"
	"github.com/gin-gonic/gin"
)

func GetByShortID(c *gin.Context) {
	shortID := c.Param("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found for given shortID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": val,
	})

}
