package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// Main manages main golang application
type Main struct {
	router *gin.Engine
}

func main() {
	m := Main{}

	// Initialize server
	if m.initServer() != nil {
		return
	}

	defer databases.Database.Close()

	c := controllers.Movie{}

	// Simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		v1.POST("/login", c.Login)
		v1.GET("/movies/list", c.ListMovies)

		// APIs need to use token string
		v1.Use(jwt.Auth(common.Config.JwtSecretPassword))
		v1.POST("/movies", c.AddMovie)
	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	m.router.Run(common.Config.Port)
}

func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize mongo database
	err = databases.Database.Init()
	if err != nil {
		return err
	}

	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()

	return nil
}
