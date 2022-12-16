package properties

import (
	"bufio"
	"connectorJIRA/pkg/logging"
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
		BindAddress       string `yaml:"bindAddress"`
		JiraUrl           string `yaml:"jiraUrl"`
		IssueInOneRequest uint   `yaml:"issueInOneRequest"`
		ThreadCount       uint   `yaml:"threadCount"`
		MaxTimeSleep      uint   `yaml:"maxTimeSleep"`
		MinTimeSleep      uint   `yaml:"minTimeSleep"`
	} `yaml:"ConnectorSettings"`
}

func GetConfig(path string) *Config {
	logger := logging.GetLogger()
	file, err := os.Open(path)
	if err != nil {
		logger.Fatal(errors.New("Error while open config file: " + err.Error()))
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		logger.Fatal(errors.New("Error while read config file: " + err.Error()))
	}

	var properties Config
	err = yaml.Unmarshal(data, &properties)
	if err != nil {
		logger.Fatal(errors.New("Error while unmarshal config file: " + err.Error()))
	}
	if properties.ProgramSettings.ThreadCount < 1 {
		logger.Fatal(errors.New("Error in config file: threadCount must be > 0"))
	}
	IssueInOneRequest := properties.ProgramSettings.IssueInOneRequest
	if IssueInOneRequest < 50 || IssueInOneRequest > 1000 {
		logger.Fatal(errors.New("Error in config file: issueInOneRequest must be >= 50 and <= 1000"))
	}

	return &properties
}
