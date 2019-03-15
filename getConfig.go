package main

import (
	"github.com/spf13/viper"
)

func getConfig() (config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/openstack-exporter/")
	viper.AddConfigPath(".")
	var configuration config

	if err := viper.ReadInConfig(); err != nil {
		return configuration, err
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, err
}
