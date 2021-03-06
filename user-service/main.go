package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"user-service/common"
	_ "user-service/docs"
	"user-service/persistent"
	"user-service/router"
)

func initServer() error {
	var err error
	// Load config file
	err = common.K8sConfig.LoadConfig()
	//err = common.K8sConfig.LocalConfig()
	if err != nil {
		return err
	}

	// Setting Gin Logger
	if common.K8sConfig.Out.EnableGinFileLog {
		//currentTime := time.Now()
		f, _ := os.Create(fmt.Sprintf("logs/gin.log"))

		if common.K8sConfig.Out.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.K8sConfig.Out.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	// Initialize User database
	err = persistent.Database.Init()
	if err != nil {
		return err
	}

	// initial Gin router
	r := router.Router{}
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
