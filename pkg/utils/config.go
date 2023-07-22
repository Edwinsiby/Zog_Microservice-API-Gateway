package utils

import "github.com/spf13/viper"

type Config struct {
	DB     string `mapstructure:"DB"`
	DSN    string `mapstructure:"DSN"`
	JWTKEY string `mapstructure:"JWTKEY"`
	INFO   SwagInfo
}

type SwagInfo struct {
	Version          string   `mapstructure:"VERSION"`
	Host             string   `mapstructure:"HOST"`
	BasePath         string   `mapstructure:"BASEPATH"`
	Schemes          []string `mapstructure:"SCHEMES"`
	Title            string   `mapstructure:"TITLE"`
	Description      string   `mapstructure:"DESCP"`
	InfoInstanceName string   `mapstructure:"INTANCENAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return Config{}, nil
}
