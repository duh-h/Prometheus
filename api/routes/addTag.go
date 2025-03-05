package routes

import (
	"encoding/json"
	"net/http"

	"github.com/duh-h/Prometheus/api/database"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag     string `json:"tag"`
}

func AddTag(c *gin.Context) {
	var tagRequest TagRequest

	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid resquest body",
		})
		return

	}

	shortID := tagRequest.ShortID
	tag := tagRequest.Tag

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found for the given shortID",
		})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		data = make(map[string]interface{})
		data["data"] = val

	}

	var tags []string

	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if strTag, ok := t.(string); ok {
				tags = append(tags, strTag)
			}
		}
	}

	for _, existingTags := range tags {
		if existingTags == tag {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "tag already exixst",
			})
			return
		}
	}

	tags = append(tags, tag)
	data["tags"] = tags

	updateData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to marshell update data",
		})
		return
	}

	err = r.Set(database.Ctx, shortID, updateData, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed update the databese",
		})
		return
	}

	c.JSON(http.StatusOK, data)

}
