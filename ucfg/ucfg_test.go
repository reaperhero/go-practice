package ucfg

import (
	"fmt"
	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/yaml"
	"os"
	"testing"
)

// go-ucfg
// ucfg is a Golang library to handle hjson, json, and yaml configuration files in your Golang project
// ufcg allows you to load yaml configuration files using . instead of indentation

type person struct {
	Age int `config:"age"`
}

type ExampleConfig struct {
	Counter int    `config:"counter" validate:"min=0, max=9"`
	Person  person `config:"config"`
}

// Defines default config option
var (
	defaultConfig = ExampleConfig{
		Counter: 8,
		Person: person{
			Age: 8,
		},
	}
)

func TestYAMLJson(t *testing.T) {
	path := "./ucfg.yaml"
	config, err := yaml.NewConfigWithFile(path, ucfg.PathSep("."))
	err = config.Unpack(&defaultConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(defaultConfig)
}
