package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

func extractStringValue(key string) string {
	panicIfNotPresent(key)
	return viper.GetString(key)
}

func panicIfNotPresent(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	}
}

func extractIntOrDefault(key string, defaultValue int) int {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return extractIntValue(key)
}

func extractIntValue(key string) int {
	panicIfNotPresent(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid Integer value", key))
	}

	return v
}
