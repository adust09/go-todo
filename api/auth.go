package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type imageraRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func imagera(c *gin.Context) {
	var req imageraRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User imageraed successfully",
		"userId":  "sample-user-id", // 実際には生成されたユーザーIDを返す
	})
}
