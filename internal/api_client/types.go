package api_client

import (
	"github.com/ivankuchin/phpipamsync/internal/config_reader"
)

type ApiClientAuthenticator interface {
	Login(cfg_original *config_reader.Config) error
	Call(req_type, urlPath, req_body string) ([]byte, error)
	Logout() error
}
