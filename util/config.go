package util

import "github.com/spf13/viper"

// Config stores all configuration of the application
// The values are read by Viper from a config file or enviroment variables
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBUrl         string `mapstructure:"DB_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
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
	return
}
