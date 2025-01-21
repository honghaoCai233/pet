package configs

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	"pet/pkg/logger"
)

type App struct {
	Addr string `yaml:"Addr"`
	Mode string `yaml:"Mode"`
}
type DB struct {
	Dialect     string `yaml:"Dialect"`
	DSN         string `yaml:"DSN"`
	MaxIdle     int    `yaml:"MaxIdle"`
	MaxActive   int    `yaml:"MaxActive"`
	MaxLifetime int    `yaml:"MaxLifetime"`
	AutoMigrate bool   `yaml:"AutoMigrate"`
}
type Redis struct {
	Addr     string `yaml:"Addr"`
	DB       int    `yaml:"DB"`
	Password string `yaml:"Password"`
}

type Config struct {
	App      App           `yaml:"App"`
	MasterDB DB            `yaml:"MasterDB"`
	SlaveDB  DB            `yaml:"SlaveDB"`
	StatsDB  DB            `yaml:"StatsDB"`
	Redis    Redis         `yaml:"Redis"`
	Log      logger.Config `yaml:"Log"`
}

func (c *Config) IsLocalOrDebugMode() bool {
	return c.IsLocalMode() || c.IsDebugMode()
}
func (c *Config) IsLocalMode() bool {
	return c.App.Mode == "local"
}
func (c *Config) IsDebugMode() bool {
	return c.App.Mode == "debug"
}

func (c *Config) IsReleaseMode() bool {
	return c.App.Mode == "release"
}

func InitConfig() (*Config, error) {
	var cfg Config
	configPath := "configs/prod.config.yaml"
	mode := os.Getenv("APP_MODE")
	if mode != "" {
		configPath = "configs/" + mode + ".config.yaml"
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err = viper.ReadConfig(file); err != nil {
		return nil, err
	}
	if err = viper.UnmarshalExact(&cfg); err != nil {
		return nil, err
	}
	return &cfg, err
}
