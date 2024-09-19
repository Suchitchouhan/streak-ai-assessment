package controller

import (
	"assessment/model"
	"assessment/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PairedNumberAPI(c *gin.Context) {
	var payload model.RequestPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := utils.FindPair(payload.Numbers, payload.Target)
	c.JSON(http.StatusOK, gin.H{"solutions": response})

}
