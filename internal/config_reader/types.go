package config_reader

type Config struct {
	Ipam_site_url string   `yaml:"ipam_site_url"`
	Ipam_app_id   string   `yaml:"ipam_app_id"`
	Ipam_app_code string   `yaml:"ipam_app_code"`
	Ipam_subnets  []string `yaml:"ipam_subnets"`
	Domain        string   `yaml:"domain"`
	Pi_hole       string   `yaml:"pi_hole"`
}

type ConfigReader interface {
	GetConfig() (*Config, error)
}
