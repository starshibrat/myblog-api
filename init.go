package main

import (
	"example/blog-service-gin/handler"
	"example/blog-service-gin/model"
	"example/blog-service-gin/repository"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	router := init_router()

	dbstore, err := init_client()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer dbstore.Disconnect()

	user_repo := repository.NewUserRepository(dbstore)
	post_repo := repository.NewPostRepository(dbstore)

	router.POST("/register", handler.Register_handler(user_repo))
	router.POST("/login", handler.Login_handler(user_repo))
	router.GET("/users", handler.Get_all_users_handler(user_repo))
	router.GET("/protected", handler.AuthenticateJwt(), handler.Protected)
	router.DELETE("deleteUser", handler.AuthenticateJwt(), handler.Delete_user_by_id(user_repo))

	router.POST("/new_post", handler.AuthenticateJwt(), handler.Create_new_post_handler(post_repo))
	router.DELETE("/post", handler.AuthenticateJwt(), handler.Delete_post_by_id_handler(post_repo))
	router.GET("/posts", handler.Get_All_Posts(post_repo))

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()

}

func init_router() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return router
}

func init_client() (model.DbStore, error) {
	clientOptions := options.Client().
		ApplyURI("mongodb://localhost:27017")

	dbstore, err := model.NewDbStore(clientOptions)

	if err != nil {
		log.Fatalf("error connecting to mongoDB: %v", err)
	}

	return dbstore, nil
}
