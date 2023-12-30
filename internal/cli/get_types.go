package cli

type IpamSubnet_list struct {
	// Code int `json:"code"`
	Subnets []Ipam_subnet_data `json:"data"`
}

type Ipam_subnet_data struct {
	Id string `json:"id"`
}

type IPAddresses struct {
	IPAddresses []IPAddress `json:"data"`
}

type IPAddress struct {
	ID                    string `json:"id"`
	SubnetID              string `json:"subnetId"`
	IP                    string `json:"ip"`
	IsGateway             string `json:"is_gateway"`
	Description           string `json:"description"`
	Hostname              string `json:"hostname"`
	Mac                   string `json:"mac"`
	Owner                 string `json:"owner"`
	Tag                   string `json:"tag"`
	DeviceID              string `json:"deviceId"`
	Location              string `json:"location"`
	Port                  string `json:"port"`
	Note                  string `json:"note"`
	LastSeen              string `json:"lastSeen"`
	ExcludePing           string `json:"excludePing"`
	PTRignore             string `json:"PTRignore"`
	Ptr                   string `json:"PTR"`
	FirewallAddressObject string `json:"firewallAddressObject"`
	EditDate              string `json:"editDate"`
	CustomerID            string `json:"customer_id"`
	ClientIdentifier      string `json:"custom_ClientID"`
}
