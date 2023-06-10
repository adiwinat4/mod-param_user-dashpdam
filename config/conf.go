package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MySQL MySQLConf
}

type MySQLConf struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func ReadConfig(filename string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(`yaml`)
	v.AddConfigPath(`.`)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found")
		} else {
			log.Println("error", err)
		}
		return nil
	}
	return v
}

func ParseConfig(v *viper.Viper) *Config {
	p := Config{}
	if err := v.Unmarshal(&p); err != nil {
		log.Println(err)
		return nil
	}
	return &p
}
