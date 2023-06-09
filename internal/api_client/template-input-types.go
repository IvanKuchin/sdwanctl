package apiclient

type TemplateInput struct {
	Entry map[string]string
}

type TemplateInputList struct {
	TemplateInput     []TemplateInput
	TemplateInputTemp []map[string]interface{} `json:"data"`
}

type RequestTemplateInput struct {
	TemplateId     string   `json:"templateId"`
	DeviceIds      []string `json:"deviceIds"`
	IsEdited       bool     `json:"isEdited"`
	IsMasterEdited bool     `json:"isMasterEdited"`
}
