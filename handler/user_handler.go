package handler

import (
	"context"
	"example/blog-service-gin/model"
	"example/blog-service-gin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_all_users_handler(repo repository.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		res, err := repo.AllUsers(context.TODO())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}

}

func Delete_user_by_id(repo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId model.UserId
		if err := c.ShouldBind(&userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cur_user_id := c.MustGet("user_id")

		if userId.Id != cur_user_id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can't delete other user"})
			return
		}

		res, err := repo.DeleteById(context.TODO(), userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, res)
	}
}
