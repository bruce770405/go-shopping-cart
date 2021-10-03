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
func (k *Local) LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	//LocalConfig = new(Local)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&LocalConfig)
	if err != nil {
		return err
	}

	// Setting Service Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   LocalConfig.c.LogFilename,
		MaxSize:    LocalConfig.c.LogMaxSize,    // megabytes after which new file is created
		MaxBackups: LocalConfig.c.LogMaxBackups, // number of backups
		MaxAge:     LocalConfig.c.LogMaxAge,     // days
	})
	log.SetLevel(log.DebugLevel)

	// log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
