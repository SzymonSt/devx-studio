package util

import "github.com/spf13/viper"

type Config struct {
	ServerPort   string
	ClientOrigin string
	DBUri        string
}

func LaodConfig() (config Config, err error) {
	viper.SetConfigName("devx")
	viper.AddConfigPath("../")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
