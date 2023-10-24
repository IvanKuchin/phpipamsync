package config_reader

var Cfg *Config

func Read() (*Config, error) {
	config_reader, err := NewConfigReader("yaml")
	if err != nil {
		return nil, err
	}

	Cfg, err = config_reader.GetConfig()
	if err != nil {
		return nil, err
	}

	return Cfg, nil
}
