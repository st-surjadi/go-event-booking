package routes

import (
	"fmt"
	"net/http"

	"example.com/go-event-booking/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Parse data is failed!",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user. Try again later!"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Create user is successful!"})
}
