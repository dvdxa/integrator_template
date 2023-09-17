package config

import "github.com/spf13/viper"

func InitConfigs() (*Config, error) {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err = viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Config struct {
	Server  Server
	Adapter Adapter
	Agent   Agent
}

type Server struct {
	Name         string
	Host         string
	Port         string
	WriteTimeout int64
	ReadTimeout  int64
}

type Adapter struct {
	URL      string
	Login    string
	Password string
	Key      string
	Account  string
	Timeout  int64
}

type Agent struct {
	Token string
	URL   string
	Key   string
}
