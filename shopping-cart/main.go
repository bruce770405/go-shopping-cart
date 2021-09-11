package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"shopping-cart/common"
	"shopping-cart/persistent"
	"shopping-cart/router"
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

	defer persistent.Database.Close()

}

func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize redis
	err = persistent.Database.Init()
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

	// initial Gin router
	r := router.Router{}
	err = r.InitRouters()
	if err != nil {
		return err
	}

	return nil
}
