package apiclient

type TemplateInput struct {
	Entry map[string]string
}

type TemplateInputTemp struct {
	Entry map[string]interface{}
}

type TemplateInputList struct {
	TemplateInput     []TemplateInput
	TemplateInputTemp []TemplateInputTemp `json:"data"`
}

type RequestTemplateInput struct {
	TemplateId     string   `json:"templateId"`
	DeviceIds      []string `json:"deviceIds"`
	IsEdited       bool     `json:"isEdited"`
	IsMasterEdited bool     `json:"isMasterEdited"`
}
