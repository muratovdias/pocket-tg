package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	RedirectURL       string
	TelegramBotURL    string `mapstructure:"bot_url"`
	DBPath            string `mapstructure:"db_file"`

	Messages Messages
}

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SavedSuccessfully string `mapstructure:"saved_successfully"`
	UnknownCommand    string `mapstructure:"unknown_command"`
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidLink  string `mapstructure:"invalid_link"`
	UnableToSave string `mapstructure:"unable_to_save"`
	Unauthorized string `mapstructure:"unauthorized"`
}

func InitConfig() (*Config, error) {

	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func parseEnv(cfg *Config) error {

	os.Setenv("TOKEN", "6268177667:AAEAYREuJ4ohhtT9tLo4Zcch2KVAQCYVWU0")
	os.Setenv("CONSUMER_KEY", "107615-645a939eb8fcde91d680be7")
	os.Setenv("REDIRECT_URL", "http://localhost:8000/")

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
