package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config nameservers to test
type Config struct {
	Concurrency  int
	Iteration    int
	Nameservers  []string
	TimeoutDial  int
	TimeoutRead  int
	TimeoutWrite int
}

// C is the global config
var C Config

func init() {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &C)
	if err != nil {
		panic(err)
	}
}
