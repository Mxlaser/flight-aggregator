package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Addr            string
	JServer1URL     string
	JServer2URL     string
	RequestTimeout  time.Duration
}

func toString(v interface{}, def string) string {
	switch t := v.(type) {
	case string:
		return t
	default:
		return fmt.Sprint(t)
	}
}

func toInt(v interface{}, def int) int {
	switch t := v.(type) {
	case int:
		return t
	case int64:
		return int(t)
	case float64:
		return int(t)
	case string:
		if n, err := strconv.Atoi(t); err == nil { return n }
	}
	return def
}

func Load() Config {
	viper.AutomaticEnv()

	viper.SetDefault("ADDR", ":8080")
	viper.SetDefault("J_SERVER1_URL", "http://j-server1:4001/flights")
	viper.SetDefault("J_SERVER2_URL", "http://j-server2:4001/flights")
	viper.SetDefault("REQUEST_TIMEOUT_SECONDS", 10)

	addr := toString(viper.Get("ADDR"), ":8080")
	js1 := toString(viper.Get("J_SERVER1_URL"), "http://j-server1:4001/flights")
	js2 := toString(viper.Get("J_SERVER2_URL"), "http://j-server2:4001/flights")
	timeoutSeconds := toInt(viper.Get("REQUEST_TIMEOUT_SECONDS"), 10)

	return Config{
		Addr:           addr,
		JServer1URL:    js1,
		JServer2URL:    js2,
		RequestTimeout: time.Duration(timeoutSeconds) * time.Second,
	}
}
