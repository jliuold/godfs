package godfs

import (
	"encoding/json"
	"errors"
	"github.com/godfs-io/godfs/logger"
	"github.com/godfs-io/godfs/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// 默认的配置文件名
var CONFIG_FILE = []string{"godfs.yaml", "godfs.yml"}

// 配置文件结构体
type Config struct {
	Godfs struct {
		Host string // 绑定的主机名或IP地址
		Port uint32 // 绑定的端口号
		Log  struct {
			Level string // 日志级别
			Path  string // 日志文件所在路径
		}
		Http struct {
			Host string
			Port int
		}
		Http2 struct {
			Host     string
			Port     int
			KeyPath  string `yaml:"key-path"`
			CertPath string `yaml:"cert-path"`
		}
		File struct {
			Storage struct {
				Path string
			}
		}
	}
}

// 全局配置文件
var CONFIG Config

// 获取默认的配置文件路径
func GetConfigPath() (string, error) {
	var configPath string

	// 判断当前目录是否存在配置文件
	for _, tmp := range CONFIG_FILE {
		path := filepath.Join(util.CurrentDir(), tmp)
		if util.IsExist(path) {
			return path, nil
		}
	}

	// 判断项目目录是否存在配置文件
	for _, tmp := range CONFIG_FILE {
		path := filepath.Join(".", tmp)
		if util.IsExist(path) {
			return path, nil
		}
	}

	logger.Error("not find config file")
	return configPath, errors.New("not find config file")
}

// 加载配置文件
func LoadConfig(file string) Config {
	bytes, _ := ioutil.ReadFile(file)
	config := Config{}
	yaml.Unmarshal(bytes, &config)
	tmp, _ := json.Marshal(config)
	logger.Debug("load config {}", string(tmp))
	return config
}

// 加载配置，初始化全局配置
func init() {
	//s, _ := GetConfigPath()
	CONFIG = LoadConfig("/Users/liujichun/workspace/godfs/godfs.yaml")
}
