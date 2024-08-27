package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) Logout(c *gin.Context) {
	// Get the user ID from the context (set by the middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Call the usecase method to perform the logout
	err := uc.userUsecase.Logout(context.Background(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfuly"})
}
