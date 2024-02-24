package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prototype_app_api/db"
	"prototype_app_api/models"
)

// / Get all markers
func MarkersGET(c *gin.Context) {
	markers, err := db.QueryMarkers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"markers": markers,
	})
}

func MarkersPOST(c *gin.Context) {
	var requestJson models.MarkerPostRequest
	if err := c.ShouldBindJSON(&requestJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowsCount, err := db.InsertMarkers(requestJson.Markers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	// create a list of markers
	c.JSON(http.StatusCreated, gin.H{"rows_affected": rowsCount, "markers": requestJson.Markers})
}

func MarkersDelete(c *gin.Context) {
	var requestJson models.MarkerDeleteRequest
	if err := c.ShouldBindJSON(&requestJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if db.DeleteMarker(requestJson.ID) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete the marker"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func MarkersPUT(c *gin.Context) {
	// update a marker
	c.JSON(http.StatusNotImplemented, gin.H{})
}
