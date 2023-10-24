package config_reader

type StaticConfigReader struct {
}

func (scr *StaticConfigReader) GetConfig() (*Config, error) {
	return &Config{}, nil
}

func NewStaticConfigReader() (ConfigReader, error) {
	return &StaticConfigReader{}, nil
}
