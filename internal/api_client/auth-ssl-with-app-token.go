package api_client

import (
	"errors"
	"log"

	"github.com/ivankuchin/phpipamsync/internal/config_reader"
)

type AuthAppCode struct {
	cfg *config_reader.Config
}

func (a *AuthAppCode) Login(cfg_original *config_reader.Config) error {
	if cfg_original == nil {
		err := errors.New("config parameter shouldn't be nil")
		log.Println(err)
		return err
	}
	a.cfg = cfg_original

	var headers http_headers = map[string]string{
		"token": a.cfg.Ipam_app_code,
	}

	_, err := sendRequestToServer("GET", a.cfg.Ipam_site_url+"/api/"+a.cfg.Ipam_app_id+"/user/", headers, "")

	if err != nil {
		log.Println("ERROR: Authentication failed:", err)
		return err
	}

	return nil
}

func (a *AuthAppCode) Call(req_type, urlPath, req_body string) ([]byte, error) {

	var headers http_headers = map[string]string{
		"token": a.cfg.Ipam_app_code,
	}

	resp, err := sendRequestToServer(req_type, urlPath, headers, req_body)
	if err != nil {
		err := errors.New("Error in call " + urlPath + " to server: " + err.Error())
		log.Println(err)
		return nil, err
	}

	return resp, nil
}

func (a *AuthAppCode) Logout() error {
	return nil
}
