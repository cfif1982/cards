package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string `yaml:"env" env:"Env" env-default:"local"`
	Server `yaml:"server"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func MustLoad() *Config {
	path := fetchConfigPath() // получаем путь к файлу конфига
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
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
