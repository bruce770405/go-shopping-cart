package router

import (
	"fmt"
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
	fmt.Println("registerRouterLinks init")
	c := controller.User{}

	// Simple group: v1
	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth", c.Authenticate)
		}

		user := v1.Group("/users")
		user.POST("", c.AddUser)
		// APIs need to be authenticated
		user.Use(jwt.Auth(common.K8sConfig.Out.JwtSecretPassword))
		{
			user.GET("/list", c.ListUsers)
			user.GET("detail/:id", c.GetUserByID)
			user.GET("/", c.GetUserByParams)
			user.DELETE(":id", c.DeleteUserByID)
			user.PATCH("", c.UpdateUser)
		}
	}

	h := controller.Health{}
	router.GET("/health", h.HealthGET)
	router.GET("/liveness", h.ReadyGET)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(common.K8sConfig.Out.Port)
	if err != nil {
		return err
	}

	return nil
}
