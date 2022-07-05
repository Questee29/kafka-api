package configs

import (
	"github.com/spf13/viper"
)

type producer struct {
}
type consumer struct {
	NAME          string   `mapstructure:"NAME"`
	KAFKA_USER    string   `mapstructure:"KAFKA_USER"`
	KAFKA_PASS    string   `mapstructure:"KAFKA_PASS"`
	KAFKA_TOPIC   string   `mapstructure:"KAFKA_TOPIC"`
	KAFKA_BROKERS []string `mapstructure:"KAFKA_BROKERS"`
}
type Config struct {
	Conumser producer `mapstructure:",squash"`
	Consumer consumer `mapstructure:",squash"`
}

func LoadConfig(configFile string, paths ...string) (Config, error) {
	config := Config{}
	viper.SetConfigName(configFile)
	viper.SetConfigType("env")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
