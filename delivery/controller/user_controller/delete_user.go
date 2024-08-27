package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bc *UserController) Deleteuser(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	isAdmin := c.GetBool("is_admin")
	if !isAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only admin can delete users"})
		return
	}

	err = bc.userUsecase.DeleteUser(context.Background(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user removed successfully",
	})
}
