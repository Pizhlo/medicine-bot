package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser    string `mapstructure:"POSTGRES_USER,required"`
	DBPass    string `mapstructure:"POSTGRES_PASSWORD,required"`
	DBName    string `mapstructure:"POSTGRES_DB,required"`
	DBHost    string `mapstructure:"POSTGRES_HOST,required"`
	DBPort    string `mapstructure:"POSTGRES_PORT,required"`
	DBAddress string
	Token     string        `mapstructure:"TOKEN,required"`
	Timeout   time.Duration `mapstructure:"TIMEOUT,required"`
}

func LoadConfig(file, path string) (*Config, error) {
	viper.SetConfigFile(file)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	conf := Config{}

	err := viper.ReadInConfig()
	if err != nil {
		return &conf, fmt.Errorf(`unable to load config: %w`, err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return &conf, fmt.Errorf(`unable to unmarshal config: %w`, err)
	}

	conf.DBAddress = fmt.Sprintf(`postgresql://%s:%s@%s:%s/%s?sslmode=disable`, conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)

	return &conf, nil
}
