package handler

import (
	"context"
	"example/blog-service-gin/model"
	"example/blog-service-gin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register_handler(repo repository.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		var registerJson model.Register
		if err := c.ShouldBindJSON(&registerJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := repo.Register(context.TODO(), &registerJson)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)
	}

}
