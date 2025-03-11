package handler

import (
	"context"
	"example/blog-service-gin/model"
	"example/blog-service-gin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create_new_post_handler(repo repository.PostRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		var postJson model.PostRequest
		if err := c.ShouldBindJSON(&postJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user_id := c.MustGet("user_id")
		new_post := model.Post{AuthorId: user_id.(string), Content: postJson.Content}

		res, err := repo.Create(context.TODO(), &new_post)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)
	}

}
