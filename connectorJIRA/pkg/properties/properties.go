package properties

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	DbSettings struct {
		DbName     string `yaml:"dbName"`
		DbPort     string `yaml:"dbPort"`
		DbHost     string `yaml:"dbHost"`
		DbUsername string `yaml:"dbUsername"`
		DbPassword string `yaml:"dbPassword"`
	} `yaml:"DBSettings"`
	ProgramSettings struct {
		ApacheUrl    string `yaml:"apacheUrl"`
		ProjectNames string `yaml:"projectNames"`
	} `yaml:"ProgramSettings"`
}

func GetConfig(path string) *Config {
	file, err := os.Open(path + "\\connectorJIRA\\config\\config.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	var properties Config
	if err := yaml.Unmarshal(data, &properties); err != nil {
		panic(err)
	}

	return &properties
}
