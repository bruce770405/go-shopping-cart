package router

import (
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"user-service/common"
	"user-service/controller"
)

// Main manages main golang application
type Router struct {
	router *gin.Engine
}

func (r *Router) InitRouters() error {
	r.router = gin.Default()
	err := registerRouterLinks(r.router)
	if err != nil {
		return err
	}

	return nil
}

func registerRouterLinks(router *gin.Engine) error {
	c := controller.User{}

	// Simple group: v1
	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth", c.Authenticate)
		}

		user := v1.Group("/users")

		// APIs need to be authenticated
		user.Use(jwt.Auth(common.Config.JwtSecretPassword))
		{
			user.POST("", c.AddUser)
			user.GET("/list", c.ListUsers)
			user.GET("detail/:id", c.GetUserByID)
			user.GET("/", c.GetUserByParams)
			user.DELETE(":id", c.DeleteUserByID)
			user.PATCH("", c.UpdateUser)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(common.Config.Port)
	if err != nil {
		return err
	}

	return nil
}
