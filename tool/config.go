package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

var _cfg *Config = nil

// ParseConfig 解析json文件
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
