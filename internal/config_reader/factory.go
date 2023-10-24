package config_reader

import "fmt"

func NewConfigReader(s string) (ConfigReader, error) {
	switch s {
	case "yaml":
		return NewYAMLConfigReader()
	case "static":
		return NewStaticConfigReader()
	}

	err := fmt.Errorf("ERROR: unknown ConfigReader type:%s", s)
	fmt.Println(err.Error())

	return nil, err
}
