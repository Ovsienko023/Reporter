package configuration

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("Not found locale .env file", err.Error())
	}
}

type Config struct {
	Api Api `yaml:"api"`
	Db  Db  `yaml:"db"`
	Tcs Tcs `yaml:"tcs"`
}

type (
	Api struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		TokenSecret string `yaml:"token_secret"`
		Doc         Doc    `yaml:"doc"`
	}

	Doc struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
)

type Db struct {
	ConnStr string `yaml:"conn_str"`
}

type Tcs struct {
	Host         string `yaml:"host"`
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

const (
	DefaultConfigPath = ""

	DefaultApiHost = "0.0.0.0"
	DefaultApiPort = "8888"

	DefaultDocHost = "127.0.0.1"
	DefaultDocPort = "8888"

	DefaultDbConnStr = "postgresql://postgres:1234@localhost:5442/postgres"
)

func NewConfig() (*Config, error) {
	cfg := &Config{
		Api{
			Host: DefaultApiHost,
			Port: DefaultApiPort,
			Doc: Doc{
				Host: DefaultDocHost,
				Port: DefaultDocPort,
			},
		},
		Db{
			ConnStr: DefaultDbConnStr,
		},
		Tcs{
			Host:         "",
			ClientId:     "",
			ClientSecret: "",
		},
	}

	if key, ok := os.LookupEnv("RP_DATABASE_CONN_STRING"); ok {
		cfg.Db.ConnStr = key
	}

	if key, ok := os.LookupEnv("RP_DOC_HOST"); ok {
		cfg.Api.Doc.Host = key
	}

	if key, ok := os.LookupEnv("RP_DOC_PORT"); ok {
		cfg.Api.Doc.Port = key
	}

	if key, ok := os.LookupEnv("RP_API_TOKEN_SECRET"); ok {
		cfg.Api.TokenSecret = key
	}

	if key, ok := os.LookupEnv("RP_TCS_HOST"); ok {
		cfg.Tcs.Host = key
	}

	if key, ok := os.LookupEnv("RP_TCS_CLIENT_ID"); ok {
		cfg.Tcs.ClientId = key
	}

	if key, ok := os.LookupEnv("RP_TCS_CLIENT_SECRET"); ok {
		cfg.Tcs.ClientSecret = key
	}

	var err error

	switch {
	case *ConfigPathFlag != DefaultConfigPath:
		err = cleanenv.ReadConfig(*ConfigPathFlag, cfg)
	case len(DefaultConfigPath) > 0:
		err = cleanenv.ReadConfig(DefaultConfigPath, cfg)
	}

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
