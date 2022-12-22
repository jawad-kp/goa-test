package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type MongoConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Collection string `mapstructure:"collection"`
}

type Config struct {
	DbConfig MongoConfig `mapstructure:"db_config"`
}

func LoadConfig(config *Config) (err error) {
	configFilePath, ok := os.LookupEnv("CONFIG_FILE_PATH")
	if !ok {
		return errors.New("CONFIG_FILE_PATH not set")
	}
	viper.SetConfigName("run")
	viper.AddConfigPath(configFilePath)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Errorf("error reading config file %s", err)
		return
	}
	log.Println("Config loaded successfully...")
	log.Println("Getting environment variables...")
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Errorf("unable to decode config to struct, %v", err)
		return
	}
	return

}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
