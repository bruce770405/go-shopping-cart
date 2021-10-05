package common

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// Config shares the global configuration
var (
	K8sConfig K8s
)

// viper import level：
// explicit call to Set: in viper use Set() default value
// flag：command value
// env
// config
// key/value store：etcd or consul
// default
// LoadConfig loads configuration from the config file
func LoadConfig() error {
	K8sConfig := new(K8s)

	v := viper.New()
	v.SetDefault(varPathToConfig, "/etc/config/config.json")
	v.SetDefault(varLogLevel, "info")
	v.AutomaticEnv()
	//v.SetConfigName("config") // default can search this file name
	//v.AddConfigPath("./config/")
	v.SetConfigFile(GetPathToConfig(v))
	err := v.ReadInConfig() // Find and read the config file
	log.WithField("path", GetPathToConfig(v)).Warn("loading config")
	if _, ok := err.(*os.PathError); ok {
		log.Warnf("no config file '%s' not found. Using default values", "config.json")
	} else if err != nil { // Handle other errors that occurred while reading the config file
		panic(fmt.Errorf("fatal error while reading the config file: %s", err))
	}

	err = v.Unmarshal(&K8sConfig.Out)

	// TODO another value need reflash when config change
	setLog()
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.WithField("file", e.Name).Warn("Config file changed")
		// TODO another value need reflash when config change
		setLog()
	})
	return nil
}

// Setting Service Logger
func setLog() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   K8sConfig.Out.Log.LogFilename,
		MaxSize:    K8sConfig.Out.Log.LogMaxSize,    // megabytes after which new file is created
		MaxBackups: K8sConfig.Out.Log.LogMaxBackups, // number of backups
		MaxAge:     K8sConfig.Out.Log.LogMaxAge,     // days
	})
	log.SetLevel(log.DebugLevel)
	// log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})
}

// GetPathToConfig returns the path to the config file
func GetPathToConfig(v *viper.Viper) string {
	return v.GetString(varPathToConfig)
}
