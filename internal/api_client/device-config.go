package apiclient

import configreader "github.com/ivankuchin/sdwanctl/internal/config-reader"

func getDeviceConfigByDevID(cfg *configreader.Config, device_id string) (string, error) {
	body, err := sendRequestToServer(cfg, "GET", cfg.Server+"/dataservice/device/config?deviceId="+device_id, "application/json", "")
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetDeviceConfig(cfg *configreader.Config, device_id string) (string, error) {
	result := ""

	config, err := getDeviceConfigByDevID(cfg, device_id)
	if err != nil {
		return result, err
	}

	return config, nil
}
