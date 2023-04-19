package apiclient

type DeviceTemplate struct {
	TemplateID     string              `json:"templateId"`
	Device         []map[string]string `json:"device"`
	IsEdited       bool                `json:"isEdited"`
	IsMasterEdited bool                `json:"isMasterEdited"`
}

type AttachFeature struct {
	DeviceTemplateList []DeviceTemplate `json:"deviceTemplateList"`
}
