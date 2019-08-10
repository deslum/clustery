package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonConfig struct {
	Database  DB     `json:"database"`
	Algorithm Engine `json:"algorithm"`
	Output    Output `json:"output"`
}

type DB struct {
	Engine string `json:"engine"`
	URI    string `json:"uri"`
}



type Engine struct {
	Name   string      `json:"name"`
	Config interface{} `json:"config"`
}

type Output struct {
	Format   string `json:"format"`
	FileName string `json:"filename"`
}

func NewJsonConfig() *JsonConfig {
	return &JsonConfig{
	}
}

func (o *JsonConfig) Fill() (*JsonConfig, error) {
	var cfg *JsonConfig
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err = json.Unmarshal(byteValue, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (o *JsonConfig) String() string {
	return fmt.Sprintf(`JSONConfig
				Database:
					Engine: %v,
					URI: %v,
				Output:
					Format: %v,
					Filename: %v`,
		o.Database.Engine,
		o.Database.URI,
		o.Output.Format,
		o.Output.FileName,
	)
}
