package apiclient

import (
	"encoding/json"
	"errors"
	"fmt"

	configreader "github.com/ivankuchin/sdwanctl/internal/config-reader"
)

func DescribeTemplateInputByDevID(cfg *configreader.Config, device_id string) (TemplateInputList, error) {
	system_devices, err := DescribeSystemDeviceByDevID(cfg, device_id)
	if err != nil {
		return TemplateInputList{}, err
	}

	var req_body RequestTemplateInput
	req_body.TemplateId = system_devices.Devices[0].TemplateID
	req_body.DeviceIds = append(req_body.DeviceIds, system_devices.Devices[0].UUID)

	req_body_json, err := json.Marshal(req_body)
	if err != nil {
		error_message := "failed to marshal request body"
		fmt.Printf("ERROR: %s (%s)\n", error_message, err.Error())
		return TemplateInputList{}, errors.New(error_message)
	}

	content, err := sendRequestToServer(cfg, "POST", cfg.Server+"/dataservice/template/device/config/input/", "application/json", string(req_body_json))
	if err != nil {
		return TemplateInputList{}, err
	}

	system_device_list := TemplateInputList{}
	if err := json.Unmarshal(content, &system_device_list); err != nil {
		error_message := "failed to unmarshal system device list"
		fmt.Printf("ERROR: %s (%s)\n", error_message, err.Error())
		return TemplateInputList{}, errors.New(error_message)
	}

	for _, temp_input_temp := range system_device_list.TemplateInputTemp {
		temp_input := TemplateInput{}
		temp_input.Entry = make(map[string]string)

		for k, v := range temp_input_temp {
			temp_input.Entry[k] = fmt.Sprintf("%v", v)
		}
		system_device_list.TemplateInput = append(system_device_list.TemplateInput, temp_input)
	}

	return system_device_list, nil
}
