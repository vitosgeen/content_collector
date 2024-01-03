package config

import (
	"content_collector/internal/apperrors"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

const (
	mongoPrefix          = "MONGODB_"
	seleniumChromePrefix = "CHROME_DRIVER_"
)

type Config struct {
	Port     string `env:"PORT,required"`
	Mongo    *MongoConfig
	Selenium *SeleniumChromeConfig
}

type SeleniumChromeConfig struct {
	Environment      string `env:"ENVIRONMENT,required"`
	LogLevel         string `env:"LOG_LEVEL,required"`
	Port             string `env:"PORT,required"`
	ChromeDriverPath string `env:"PATH,required"`
	ChromeDriverPort int    `env:"PORT,required"`
}

type MongoConfig struct {
	MongoDbUri  string `env:"URI,required"`
	MongoDBHost string `env:"HOST,required"`
	MongoDBPort string `env:"PORT,required"`
	MongoDbUser string `env:"USER,required"`
	MongoDbPass string `env:"PASS,required"`
	MongoDbName string `env:"NAME,required"`
}

func NewConfig(envStr string) (*Config, error) {
	err := godotenv.Load(envStr)
	if err != nil {
		return nil, apperrors.EnvConfigLoadError.AppendMessage(err)
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return cfg, apperrors.EnvConfigParseError.AppendMessage(err)
	}

	err = cfg.AddMongoConfig(cfg)
	if err != nil {
		return cfg, err
	}

	err = cfg.AddSeleniumChromeConfig(cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (c *Config) AddMongoConfig(cfg *Config) error {
	mongoCfg := &MongoConfig{}
	opts := env.Options{
		Prefix: mongoPrefix,
	}
	if err := env.ParseWithOptions(mongoCfg, opts); err != nil {
		return apperrors.EnvConfigMongoParseError.AppendMessage(err)
	}
	cfg.Mongo = mongoCfg

	return nil
}

func (c *Config) AddSeleniumChromeConfig(cfg *Config) error {
	seleniumCfg := &SeleniumChromeConfig{}
	opts := env.Options{
		Prefix: seleniumChromePrefix,
	}
	if err := env.ParseWithOptions(seleniumCfg, opts); err != nil {
		return apperrors.EnvConfigSeleniumChromeParseError.AppendMessage(err)
	}
	cfg.Selenium = seleniumCfg

	return nil
}

func (c *Config) GetMongoConfig() *MongoConfig {
	return c.Mongo
}

func (c *Config) GetSeleniumChromeConfig() *SeleniumChromeConfig {
	return c.Selenium
}
