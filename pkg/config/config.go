package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	RedirectURL       string
	TelegramBotURL    string
}

func InitConfig() (*Config, error) {
	var cfg Config
	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func parseEnv(cfg *Config) error {

	if err := viper.BindEnv("TOKEN"); err != nil {
		return err
	}
	if err := viper.BindEnv("CONSUMER_KEY"); err != nil {
		return err
	}
	if err := viper.BindEnv("REDIRECT_URL"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("TOKEN")
	cfg.PocketConsumerKey = viper.GetString("CONSUMER_KEY")
	cfg.RedirectURL = viper.GetString("REDIRECT_URL")
	return nil
}
