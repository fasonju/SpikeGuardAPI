package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prototype_app_api/db"
	"prototype_app_api/models"
)

func ReportPUT(c *gin.Context) {
	var requestJson models.ReportPutRequest
	if err := c.ShouldBindJSON(&requestJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.InsertReport(requestJson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
