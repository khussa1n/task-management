package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	HTTP  ServerConfig `yaml:"http"`
	DB    DBConfig     `yaml:"db"`
	Token TokenConfig  `yaml:"token"`
}

type TokenConfig struct {
	SecretKey  string        `yaml:"token_secret_key"`
	TimeToLive time.Duration `yaml:"time_to_live"`
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
}

type DBConfig struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	DBName           string `yaml:"db_name"`
	Username         string `yaml:"username"`
	MigrationPath    string `json:"migration_path"`
	MigrationVersion uint   `json:"migration_version"`
	Password         string `yaml:"password"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)

	return cfg, nil
}
