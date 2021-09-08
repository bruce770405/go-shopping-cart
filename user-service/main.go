package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
	"user-service/common"
	_ "user-service/docs"
	"user-service/persistent"
	"user-service/routers"
)

func initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		currentTime := time.Now()
		f, _ := os.Create(fmt.Sprintf("logs/log-%s.log", currentTime.Format("YYYY-MM-DD")))

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

	// Initialize User database
	err = persistent.Database.Init()
	if err != nil {
		return err
	}

	// initial Gin routers
	r := routers.Router{}
	err = r.InitRouters()
	if err != nil {
		return err
	}

	return nil
}

// @title UserManagement Service API Document
// @version 1.0
// @description List APIs of UserManagement Service
// @termsOfService http://swagger.io/terms/

// @host 107.113.53.47:8808
// @BasePath /api/v1
func main() {

	// Initialize server
	error := initServer()
	if error != nil {
		return
	}

	// close connection after server down
	defer persistent.Database.Close()
}
