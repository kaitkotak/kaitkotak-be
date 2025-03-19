package config

import (
	"fmt"

	filehelper "github.com/kaitkotak-be/internal/shared/file-helper"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var Value = koanf.New(".")

func LoadConfig() {
	configPath, err := filehelper.GetAbsolutePath("configs/config.yaml")
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	err = Value.Load(file.Provider(configPath), yaml.Parser())
	if err != nil {
		fmt.Println("Error loading config file:", err)
	}
	Value.Load(file.Provider("./configs/config.yml"), yaml.Parser())
}
