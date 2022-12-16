package app

import (
	"bufio"
	"gopkg.in/yaml.v3"
	"io"
	"log"
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
		BindAddress       string `yaml:"bindAddress"`
		JiraUrl           string `yaml:"jiraUrl"`
		IssueInOneRequest uint   `yaml:"issueInOneRequest"`
		ThreadCount       uint   `yaml:"threadCount"`
		MaxTimeSleep      uint   `yaml:"maxTimeSleep"`
		MinTimeSleep      uint   `yaml:"minTimeSleep"`
	} `yaml:"ProgramSettings"`
	Backend struct {
		Port       uint   `yaml:"port"`
		UpdateTime uint64 `yaml:"updateTime"`
	} `yaml:"Backend"`
}

func NewConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	var properties Config
	err = yaml.Unmarshal(data, &properties)
	if err != nil {
		log.Fatal(err)
	}
	return &properties
}
