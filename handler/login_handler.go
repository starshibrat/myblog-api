package handler

import (
	"context"
	"example/blog-service-gin/model"
	"example/blog-service-gin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login_handler(repo repository.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		var loginJson model.Login
		if err := c.ShouldBindJSON(&loginJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := repo.Login(context.TODO(), &loginJson)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}

}
