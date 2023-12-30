package api_client

type IPList struct {
	IPs []IP `json:"data"`
}

type IP struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Mac      string `json:"mac"`
	ClientID string `json:"custom_ClientID"`
}
