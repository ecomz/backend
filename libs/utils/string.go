package utils

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func checkKey(key string) {
	_, envExist := os.LookupEnv(key)

	if !viper.IsSet(key) && !envExist {
		panic("missing required key: " + key)
	}
}

func getString(key string) string {
	if val, exist := os.LookupEnv(key); exist {
		return val
	}
	return viper.GetString(key)
}

func GetStringOrPanic(key string) string {
	checkKey(key)
	return getString(key)
}

func GetIntOrDefault(key string, def int) int {
	val := getString(key)
	if intVal, err := strconv.Atoi(val); err != nil {
		return intVal
	}
	return def
}

func GetIntOrPanic(key string) int {
	checkKey(key)
	return GetIntOrDefault(key, 0)
}
