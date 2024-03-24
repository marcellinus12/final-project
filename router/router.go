package router

import (
	"go-gorm/controllers"
	"go-gorm/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", controllers.UserRegister)
		authGroup.POST("/login", controllers.UserLoginn)
	}

	userGroup := r.Group("/users")
	{
		userGroup.Use(middlewares.Authentication())
		userGroup.PUT("/:userID", middlewares.UserAuthorization(), controllers.UpdateUser)
		userGroup.DELETE("/:userID", middlewares.UserAuthorization(), controllers.DeleteUser)
	}

	photoGroup := r.Group("/photo")
	{
		photoGroup.Use(middlewares.Authentication())
		photoGroup.GET("/", controllers.GetAllPhotos)
		photoGroup.POST("/", controllers.CreatePhoto)
		photoGroup.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoGroup.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentGroup := r.Group("/comment")
	{
		commentGroup.Use(middlewares.Authentication())
		commentGroup.GET("/", controllers.GetAllComments)
		commentGroup.POST("/", controllers.CreateComment)
		commentGroup.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentGroup.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaGroup := r.Group("/socialMedia")
	{
		socialMediaGroup.Use(middlewares.Authentication())
		socialMediaGroup.GET("/", controllers.GetAllSocialMedia)
		socialMediaGroup.POST("/", controllers.CreateSocialMedia)
		socialMediaGroup.PUT("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaGroup.DELETE("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
