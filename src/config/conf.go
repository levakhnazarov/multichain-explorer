package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

type (
	Config struct {
		*viper.Viper
	}
)

const (
	DefaultEnvironment = "prod"
	envName            = "ENV"
	configPathName     = "CONFIG_PATH"
)

var (
	env        = ""
	configPath = ""

	path string
)

func init() {
	env = os.Getenv(envName)
	configPath = os.Getenv(configPathName)
	flag.StringVar(&path, "c", "", "path to config file")
	flag.Parse()
}

//TODO singleton
func NewConfig(envName string) *Config {
	v := viper.New()

	if len(env) != 0 {
		envName = env
	}

	v.SetConfigName(envName)

	if len(configPath) != 0 {
		v.AddConfigPath(configPath)
	}

	v.AddConfigPath("../../")

	v.AddConfigPath(path)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	if err != nil {
		log.Fatal("Error load config ", "error", err.Error())
	}

	return &Config{v}
}
