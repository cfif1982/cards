package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const cfgTypeYAML = "YAML"
const cfgTypeENV = "ENV"
const environmentPrefix = "CARD"

type Config struct {
	Env      string `yaml:"env" env:"Env" env-default:"local"`
	Server   `yaml:"server"`
	Database `yaml:"database"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"table"`
}

func MustLoad() *Config {
	var cfg Config

	// указываем откуда брать настройки
	var currentCfgType = cfgTypeENV

	// если настройки берутся из yaml файла
	if currentCfgType == cfgTypeYAML {
		path := fetchConfigPath() // получаем путь к файлу конфига
		if path == "" {
			panic("config path is empty")
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			panic("config file does not exist: " + path)
		}

		if err := cleanenv.ReadConfig(path, &cfg); err != nil {
			panic("failed to read config: " + err.Error())
		}

		return &cfg
		// если настройки берутся из env
	} else {
		// загружаем файл настроек для среды окружения.
		// считываем этот файл и инициализируем соответсвтующие переменные среды окружения
		err := godotenv.Load("config.env")
		if err != nil {
			fmt.Println("error while loading config.env file")
		}

		// cfg := &Config{}
		if err = envconfig.Process(environmentPrefix, &cfg); err != nil {
			panic("error while loading environment")
		}

		return &cfg
	}
}

// Получаем путь к файлу config из флага или среды окружения
// Приоритет: флаг > среда > по-умолчанию
// значение по-умолчанию: ""
// запуск с указанием в среде окружения: CONFIG_PATH=./path/file.yaml cards
// запуск с указанием флага: cards --config=./path/file.yaml
func fetchConfigPath() string {
	var res string

	// сначала проверяем флаг
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	// есом флага нет, то берем из среды окружения
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
