package apiclient

type Device struct {
	DeviceID            string   `json:"deviceId"`
	SystemIP            string   `json:"system-ip"`
	HostName            string   `json:"host-name"`
	Reachability        string   `json:"reachability"`
	Status              string   `json:"status"`
	Personality         string   `json:"personality"`
	DeviceType          string   `json:"device-type"`
	Timezone            string   `json:"timezone"`
	DeviceGroups        []string `json:"device-groups"`
	Lastupdated         int64    `json:"lastupdated"`
	DomainID            string   `json:"domain-id,omitempty"`
	BoardSerial         string   `json:"board-serial"`
	CertificateValidity string   `json:"certificate-validity"`
	MaxControllers      string   `json:"max-controllers,omitempty"`
	UUID                string   `json:"uuid"`
	ControlConnections  string   `json:"controlConnections"`
	DeviceModel         string   `json:"device-model"`
	Version             string   `json:"version"`
	ConnectedVManages   []string `json:"connectedVManages"`
	SiteID              string   `json:"site-id"`
	Latitude            string   `json:"latitude"`
	Longitude           string   `json:"longitude"`
	IsDeviceGeoData     bool     `json:"isDeviceGeoData"`
	Platform            string   `json:"platform"`
	UptimeDate          int64    `json:"uptime-date"`
	StatusOrder         int      `json:"statusOrder"`
	DeviceOs            string   `json:"device-os"`
	Validity            string   `json:"validity"`
	State               string   `json:"state"`
	StateDescription    string   `json:"state_description"`
	ModelSku            string   `json:"model_sku"`
	LocalSystemIP       string   `json:"local-system-ip"`
	TotalCPUCount       string   `json:"total_cpu_count"`
	TestbedMode         bool     `json:"testbed_mode"`
	LayoutLevel         int      `json:"layoutLevel"`
	OmpPeers            string   `json:"ompPeers,omitempty"`
	LinuxCPUCount       string   `json:"linux_cpu_count,omitempty"`
	BfdSessionsUp       int      `json:"bfdSessionsUp,omitempty"`
	BfdSessions         string   `json:"bfdSessions,omitempty"`
}

// type Device struct{}

type DeviceList struct {
	Devices []Device `json:"data"`
}

type SystemDevice struct {
	DeviceType                 string   `json:"deviceType"`
	SerialNumber               string   `json:"serialNumber"`
	NcsDeviceName              string   `json:"ncsDeviceName"`
	ConfigStatusMessage        string   `json:"configStatusMessage"`
	TemplateApplyLog           []string `json:"templateApplyLog"`
	UUID                       string   `json:"uuid"`
	ManagementSystemIP         string   `json:"managementSystemIP"`
	TemplateStatus             string   `json:"templateStatus"`
	ChasisNumber               string   `json:"chasisNumber"`
	ConfigStatusMessageDetails string   `json:"configStatusMessageDetails"`
	ConfigOperationMode        string   `json:"configOperationMode"`
	DeviceModel                string   `json:"deviceModel"`
	DeviceState                string   `json:"deviceState"`
	Validity                   string   `json:"validity"`
	PlatformFamily             string   `json:"platformFamily"`
	VedgeCertificateState      string   `json:"vedgeCertificateState"`
	RootCertHash               string   `json:"rootCertHash"`
	DeviceIP                   string   `json:"deviceIP"`
	Personality                string   `json:"personality"`
	UploadSource               string   `json:"uploadSource"`
	SubjectSerialNumber        string   `json:"subjectSerialNumber"`
	LocalSystemIP              string   `json:"local-system-ip"`
	SystemIP                   string   `json:"system-ip"`
	ModelSku                   string   `json:"model_sku"`
	SiteID                     string   `json:"site-id"`
	HostName                   string   `json:"host-name"`
	Version                    string   `json:"version"`
	Vbond                      string   `json:"vbond"`
	VmanageConnectionState     string   `json:"vmanageConnectionState"`
	Lastupdated                int64    `json:"lastupdated"`
	Reachability               string   `json:"reachability"`
	UptimeDate                 int64    `json:"uptime-date"`
	DefaultVersion             string   `json:"defaultVersion"`
	AvailableVersions          []string `json:"availableVersions"`
	Template                   string   `json:"template"`
	TemplateID                 string   `json:"templateId"`
	LifeCycleRequired          bool     `json:"lifeCycleRequired"`
	ExpirationDate             string   `json:"expirationDate"`
	HardwareCertSerialNumber   string   `json:"hardwareCertSerialNumber"`
	DraftMode                  string   `json:"draftMode"`
}

type SystemDeviceList struct {
	Devices []SystemDevice `json:"data"`
}
