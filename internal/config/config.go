package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"githib.com/igomonov88/game-service/internal/database"
	"githib.com/igomonov88/game-service/internal/server"
)

type Config struct {
	Server      server.Config
	Database    database.Config
	ServiceName string
}

func Must(config *Config, err error) Config {
	if err != nil {
		panic(err)
	}

	return *config
}

func ReadConfig() (*Config, error) {
	pflag.Parse()

	vp := viper.New()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vp.SetConfigType("yml")
	vp.SetConfigName("config")
	vp.AddConfigPath("etc")
	vp.AutomaticEnv()
	vp.SetEnvPrefix("game-service")

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := vp.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
