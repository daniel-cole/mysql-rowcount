package config

import (
	"github.com/daniel-cole/mysql-rowcount/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type RowCountConfig struct {
	User                 string   `yaml:"user"`
	Passwd               string   `yaml:"password"`
	Addr                 string   `yaml:"address"`
	Net                  string   `yaml:"net"`
	AllowNativePasswords bool     `yaml:"nativepasswords"`
	DatabasesToIgnore    []string `yaml:"databasestoignore"`
	DatabasesToInclude   []string `yaml:"databasestoinclude"`
	TablesToIgnore       []string `yaml:"tablestoignore"`
	MaxWorkers           int      `yaml:"maxworkers"`
	OutputFile           string   `yaml:"outputfile"`
}

func LoadConfig() *RowCountConfig {

	configFile, err := util.GetEnv("ROWCOUNT_CONFIG", "rowcount-config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	var config *RowCountConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
