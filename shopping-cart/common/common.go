package common

type Redis struct {
	Host     string `mapstructure:"localhost:6379"`
	Password string `mapstructure:"redisPassword"`
}

type DataSource struct {
	MgAddrs      string `mapstructure:"mgAddrs"`
	MgDbName     string `mapstructure:"mgDbName"`
	MgDbUsername string `mapstructure:"mgDbUsername"`
	MgDbPassword string `mapstructure:"mgDbPassword"`
}

type Logging struct {
	LogLevel      string `mapstructure:"loglevel" validate:"required,oneof=debug info warn error"`
	LogFilename   string `mapstructure:"logFilename"`
	LogMaxSize    int    `mapstructure:"logMaxSize"`
	LogMaxBackups int    `mapstructure:"logMaxBackups"`
	LogMaxAge     int    `mapstructure:"logMaxAge"`
}

// Configuration stores setting values
type Configuration struct {
	Port                string     `mapstructure:"port"`
	EnableGinConsoleLog bool       `mapstructure:"enableGinConsoleLog"`
	EnableGinFileLog    bool       `mapstructure:"enableGinFileLog"`
	Log                 Logging    `mapstructure:"logging" validate:"required"`
	Db                  DataSource `mapstructure:"datasource" validate:"required"`
	Redis               Redis      `mapstructure:"redis" validate:"required"`
	JwtSecretPassword   string     `mapstructure:"jwtSecretPassword"`
	Issuer              string     `mapstructure:"issuer"`
}

const (
	varLogLevel     = "log.level"
	varPathToConfig = "config.file"
)

type Config struct {
	Out Configuration
}

type Local struct {
	Out Configuration
}

type K8s struct {
	Out Configuration
}

// COLLECTIONs of the database table
const (
	ColShoppingCart = "shoppingCart"
)
