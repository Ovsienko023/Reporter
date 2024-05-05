package configuration

import (
	"github.com/BurntSushi/toml"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"path/filepath"
)

//const DefaultConfigPath = ""

type Config struct {
	Api    Api    `yaml:"api"`
	Db     Db     `yaml:"db"`
	Tls    Tls    `toml:"tls"`
	Static Static `toml:"static"`
	Doc    Doc    `toml:"doc"`
	Tcs    Tcs    `yaml:"tcs"`
}

type (
	Api struct {
		Host        string `toml:"host" env:"WEB_API_HOST" env-default:"127.0.0.1"`
		Port        string `toml:"port" env:"WEB_API_PORT" env-default:"8080"`
		TokenSecret string `toml:"token_secret" env:"WEB_API_TOKEN_SECRET"`
	}

	Db struct {
		ConnStr string `toml:"conn_str" env:"WEB_API_DB_CONN" env-default:"postgresql://postgres:1234@localhost:5442/postgres"`
	}

	Tls struct {
		Enable       bool   `toml:"enable" env:"WEB_API_TLS_ENABLE" env-default:"false"`
		CertFilePath string `toml:"cert_file_path" env:"WEB_API_TLS_CERT_PATH"`
		KeyFilePath  string `toml:"key_file_path" env:"WEB_API_TLS_KEY_PATH"`
	}

	Static struct {
		FilesPath string `toml:"files_path" env:"WEB_API_STATIC_PATH" env-default:"web"`
	}

	Doc struct {
		Host string `toml:"host" env:"WEB_API_DOC_HOST" env-default:"127.0.0.1"`
		Port string `toml:"port" env:"WEB_API_DOC_PORT" env-default:"8888"`
	}

	Tcs struct {
		Host         string `toml:"host" env:"WEB_API_AUTH_TCS_HOST"`
		ClientId     string `toml:"client_id" env:"WEB_API_AUTH_TCS_CLIENT_ID"`
		ClientSecret string `toml:"client_secret" env:"WEB_API_AUTH_TCS_CLIENT_SECRET"`
	}
)

func New() (*Config, error) {
	var cfg Config

	var configFullPath string
	rawConfFullPath := filepath.Join(filepath.Clean(*ConfigPathFlag), "web-api.toml")

	if ConfigPathFlag != nil && *ConfigPathFlag != "" && isExistConfig(rawConfFullPath) {
		configFullPath = rawConfFullPath
	} else {
		configFullPath = DefaultConfigPath
	}

	if configFullPath != "" {
		if err := cleanenv.ReadConfig(filepath.Clean(*ConfigPathFlag), &cfg); err != nil {
			log.Println("Failed to read config file:", *ConfigPathFlag, err)
		}
	} else {
		log.Println("Configuration file not found!")
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Println("Failed to read env:", *ConfigPathFlag, err)
	}

	// Создание файла конфигурации на основе собранного конфига
	cfg.createToml(configFullPath)

	return &cfg, nil
}

func (c *Config) createToml(filePath string) {
	if filePath != "" {

	}

	f, err := os.Create("web-api.toml")
	if err != nil {
		// failed to create/open the file
		log.Println(err.Error())
	}
	if err := toml.NewEncoder(f).Encode(c); err != nil {
		// failed to encode
		log.Println(err.Error())
	}

	if err := f.Close(); err != nil {
		log.Println(err.Error())
	}
}

func isExistConfig(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return true
}
