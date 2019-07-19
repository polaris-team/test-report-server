package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var conf Config = Config{
	Viper:    viper.New(),
	Mysql:    nil,
	Redis:    nil,
	Mail:     nil,
	Server:   nil,
	DingTalk: nil,
	Parameters: nil,
}

type Config struct {
	Viper    *viper.Viper
	Mysql    *MysqlConfig
	Redis    *RedisConfig
	Mail     *MailConfig
	Server   *ServerConfig
	DingTalk *DingTalkSDKConfig
	Parameters *map[string]*string
}

type MysqlConfig struct {
	Host     string
	Port     int
	Usr      string
	Pwd      string
	Database string
}

type RedisConfig struct {
	Host           string
	Port           int
	Pwd            string
	Database       int
	MaxIdle        int
	MaxActive      int
	MaxIdleTimeout int
}

type MailConfig struct {
	Usr  string
	Pwd  string
	Host string
	Port int
}

type ServerConfig struct {
	Port int
	Name string
}

type DingTalkSDKConfig struct {
	SuiteKey    string
	SuiteSecret string
}

func GetMysqlConfig() *MysqlConfig {
	return conf.Mysql
}

func GetRedisConfig() *RedisConfig {
	return conf.Redis
}

func GetConfig() Config {
	return conf
}

func GetMailConfig() *MailConfig {
	return conf.Mail
}

func GetServerConfig() *ServerConfig {
	return conf.Server
}

func GetDingTalkSdkConfig() *DingTalkSDKConfig {
	return conf.DingTalk
}

func GetParameters() *map[string]*string{
	return conf.Parameters
}

func GetParameter(key string) string{
	key = strings.ToLower(key)
	if conf.Parameters == nil{
		panic(errors.New("Parameters configuration is nil!"))
	}
	ps := *conf.Parameters
	if ps[key] == nil{
		panic(errors.Errorf("Parameter %s Not configured!", key))
	}
	return *ps[key]
}

func LoadConfig(dir string, config string) error {
	return LoadEnvConfig(dir, config, "")
}

func LoadEnvConfig(dir string, config string, env string) error {
	if env != "" {
		config += env
	}
	conf.Viper.SetConfigName(config)
	conf.Viper.AddConfigPath(dir)
	conf.Viper.SetConfigType("yaml")
	if err := conf.Viper.ReadInConfig(); err != nil {
		return err
	}
	if err := conf.Viper.Unmarshal(&conf); err != nil {
		return err
	}
	return nil
}
