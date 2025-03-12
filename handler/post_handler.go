package handler

import (
	"context"
	"example/blog-service-gin/model"
	"example/blog-service-gin/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Create_new_post_handler(repo repository.PostRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		var postJson model.PostRequest
		if err := c.ShouldBindJSON(&postJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user_id, err := bson.ObjectIDFromHex(c.MustGet("user_id").(string))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		new_post := model.Post{AuthorId: user_id, Content: postJson.Content, CreatedAt: time.Now()}

		res, err := repo.Create(context.TODO(), &new_post)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)
	}

}

func Delete_post_by_id_handler(repo repository.PostRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		var postId model.PostId
		if err := c.ShouldBind(&postId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user_id := c.MustGet("user_id").(string)

		res, err := repo.DeleteById(context.TODO(), &postId, user_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, res)
	}

}

func Get_All_Posts(repo repository.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := repo.GetAllPost(context.TODO())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
