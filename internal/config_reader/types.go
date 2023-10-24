package config_reader

type Config struct {
	Server      string
	BearerToken string
	XSRFToken   string
	Login       string
	Password    string
}

type Config_file struct {
	Vmanages        []Vmanage_list `yaml:"vmanages"`
	Users           []User_list    `yaml:"users"`
	Contexts        []Context_list `yaml:"contexts"`
	Current_context string         `yaml:"current-context"`
}

type Vmanage_list struct {
	Name        string `yaml:"name"`
	Server      string `yaml:"server"`
	BearerToken string `yaml:"bearer-token"`
}

type User_list struct {
	Name     string `yaml:"name"`
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
}

type Context_list struct {
	Name    string `yaml:"name"`
	User    string `yaml:"user"`
	Vmanage string `yaml:"vmanage"`
}

type ConfigReader interface {
	GetConfig() (*Config, error)
}
