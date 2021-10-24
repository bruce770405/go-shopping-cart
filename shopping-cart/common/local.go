package common

import (
	"encoding/json"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"os"
)

// Config shares the global configuration
var (
	LocalConfig Local
)

//// LoadConfig loads configuration from the config file
func (local *Local) LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&local)
	if err != nil {
		return err
	}

	// Setting Service Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   LocalConfig.Out.Log.LogFilename,
		MaxSize:    LocalConfig.Out.Log.LogMaxSize,    // megabytes after which new file is created
		MaxBackups: LocalConfig.Out.Log.LogMaxBackups, // number of backups
		MaxAge:     LocalConfig.Out.Log.LogMaxAge,     // days
	})
	log.SetLevel(log.DebugLevel)

	// log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
