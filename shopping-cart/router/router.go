package router

import (
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"shopping-cart/common"
	"shopping-cart/controller"
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
	c := controller.Shopping{}

	// Simple group: v1
	v1 := router.Group("/api/v1")
	{
		// APIs need to use token string
		v1.Use(jwt.Auth(common.Config.JwtSecretPassword))
		{
			v1.POST("/cart/add", c.AddProdInCart)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(common.Config.Port)

	err := router.Run(common.Config.Port)
	if err != nil {
		return err
	}

	return nil
}

