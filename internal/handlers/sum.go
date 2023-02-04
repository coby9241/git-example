package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SumInput struct {
	FirstNumber  int `json:"first_number" binding:"required"`
	SecondNumber int `json:"second_number" binding:"required"`
}

func sumNumbers(first, second int) int {
	return first + second
}

func SumHandler(c *gin.Context) {
	var input SumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorCode": "ApiMissingRequiredFieldException",
			"message":   "One or more of the following required fields were missing: first_number, second_number",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result": gin.H{
			"sum": sumNumbers(input.FirstNumber, input.SecondNumber),
		},
	})
}
