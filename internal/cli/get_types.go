package cli

type IpamSubnet_list struct {
	// Code int `json:"code"`
	Subnets []Ipam_subnet_data `json:"data"`
}

type Ipam_subnet_data struct {
	Id int `json:"id"`
}

type IPAddresses struct {
	IPAddresses []IPAddress `json:"data"`
}

type IPAddress struct {
	ID                    int    `json:"id"`
	SubnetID              int    `json:"subnetId"`
	IP                    string `json:"ip"`
	IsGateway             int    `json:"is_gateway"`
	Description           string `json:"description"`
	Hostname              string `json:"hostname"`
	Mac                   string `json:"mac"`
	Owner                 string `json:"owner"`
	Tag                   int    `json:"tag"`
	DeviceID              int    `json:"deviceId"`
	Location              string `json:"location"`
	Port                  string `json:"port"`
	Note                  string `json:"note"`
	LastSeen              string `json:"lastSeen"`
	ExcludePing           int    `json:"excludePing"`
	PTRignore             int    `json:"PTRignore"`
	Ptr                   string `json:"PTR"`
	FirewallAddressObject string `json:"firewallAddressObject"`
	EditDate              string `json:"editDate"`
	CustomerID            string `json:"customer_id"`
	ClientIdentifier      string `json:"custom_ClientID"`
}
