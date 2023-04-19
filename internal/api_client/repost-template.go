package apiclient

import (
	"encoding/json"
	"errors"
	"fmt"

	configreader "github.com/ivankuchin/sdwanctl/internal/config-reader"
)

func RepostTemplate(cfg *configreader.Config, template_input TemplateInput) error {
	var attach_feature AttachFeature
	attach_feature.DeviceTemplateList = make([]DeviceTemplate, 1)
	attach_feature.DeviceTemplateList[0].Device = make([]map[string]string, 1)

	attach_feature.DeviceTemplateList[0].Device[0] = template_input.Entry
	attach_feature.DeviceTemplateList[0].IsEdited = false
	attach_feature.DeviceTemplateList[0].IsMasterEdited = false
	attach_feature.DeviceTemplateList[0].TemplateID = template_input.Entry["csv-templateId"]

	attach_feature_json, err := json.Marshal(attach_feature)
	if err != nil {
		error_message := "failed to marshal attach feature"
		fmt.Printf("ERROR: %s (%s)\n", error_message, err.Error())
		return errors.New(error_message)
	}
	// fmt.Printf("DEBUG: attach_feature_json: %s\n", attach_feature_json)

	_, err = sendRequestToServer(cfg, "POST", cfg.Server+"/dataservice/template/device/config/attachfeature", "application/json", string(attach_feature_json))
	if err != nil {
		return err
	}

	return nil
}
