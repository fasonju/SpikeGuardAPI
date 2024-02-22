package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prototype_app_api/db"
)

// / Get all markers
func MarkersGET(c *gin.Context) {
	markers, err := db.QueryMarkers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNotImplemented, gin.H{
		"markers": markers,
	})
}

func MarkersPOST(c *gin.Context) {
	// create a marker
	c.JSON(http.StatusNotImplemented, gin.H{})
}

func MarkersDelete(c *gin.Context) {
	// delete a marker
	c.JSON(http.StatusNotImplemented, gin.H{})
}

func MarkersPUT(c *gin.Context) {
	// update a marker
	c.JSON(http.StatusNotImplemented, gin.H{})
}
