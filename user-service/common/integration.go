package common

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// Config shares the global configuration
var (
	IntegrateConfig Config
)

// viper import level：
// explicit call to Set: in viper use Set() default value
// flag：command value
// env
// config
// key/value store：etcd or consul
// default
// LoadConfig loads configuration from the config file
func (config *Config) LoadConfig() error {
	err := LocalConfig.LoadConfig()
	if _, ok := err.(*os.PathError); ok {
		log.Warnf("no config file '%s' not found. Using default values", "config.json")
	} else if err != nil { // Handle other errors that occurred while reading the config file
		log.Error("fatal error while reading the config file: %s", err)
	}

	err = K8sConfig.LoadConfig()
	if _, ok := err.(*os.PathError); ok {
		log.Warnf("no config file '%s' not found. Using default values", "config.json")
	} else if err != nil { // Handle other errors that occurred while reading the config file
		log.Error("fatal error while reading the config file: %s", err)
	}

	config.Out.MgAddrs = K8sConfig.Out.MgAddrs
	config.Out.MgDbName = K8sConfig.Out.MgDbName
	config.Out.MgDbPassword = K8sConfig.Out.MgDbPassword
	config.Out.MgDbUsername = K8sConfig.Out.MgDbUsername
	config.Out.Log = K8sConfig.Out.Log
	return nil
}
