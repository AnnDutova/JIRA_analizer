package properties

import (
	"bufio"
	"errors"
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
		ApacheUrl         string `yaml:"apacheUrl"`
		IssueInOneRequest uint   `yaml:"issueInOneRequest"`
		ThreadCount       uint   `yaml:"threadCount"`
		MaxTimeSleep      uint   `yaml:"maxTimeSleep"`
		MinTimeSleep      uint   `yaml:"minTimeSleep"`
	} `yaml:"ProgramSettings"`
}

func GetConfig(path string) (*Config, error) {
	file, err := os.Open(path + "\\connectorJIRA\\config\\config.yaml")
	if err != nil {
		return nil, errors.New("Error while open config file: " + err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.New("Error while read config file: " + err.Error())
	}

	var properties Config
	err = yaml.Unmarshal(data, &properties)
	if err != nil {
		return nil, errors.New("Error while unmarshal config file: " + err.Error())
	}
	if properties.ProgramSettings.ThreadCount < 1 {
		return nil, errors.New("Error in config file: threadCount must be > 0")
	}
	IssueInOneRequest := properties.ProgramSettings.IssueInOneRequest
	if IssueInOneRequest < 50 || IssueInOneRequest > 1000 {
		return nil, errors.New("Error in config file: issueInOneRequest must be >= 50 and <= 1000")
	}

	return &properties, nil
}
