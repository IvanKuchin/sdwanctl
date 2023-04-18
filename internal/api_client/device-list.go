package apiclient

import (
	"encoding/json"
	"errors"
	"fmt"

	configreader "github.com/ivankuchin/sdwanctl/internal/config-reader"
)

func getDeviceListFromServer(cfg *configreader.Config) (DeviceList, error) {

	// Set up the HTTP client
	// Set any headers you need to include
	// Add the bearer token(bt) to the request
	// Send the request
	// Read the response body
	content, err := sendRequestToServer(cfg, "GET", cfg.Server+"/dataservice/device", "application/json", "")
	if err != nil {
		return DeviceList{}, err
	}

	device_list := DeviceList{}
	if err := json.Unmarshal(content, &device_list); err != nil {
		error_message := "failed to unmarshal device list"
		fmt.Printf("ERROR: %s (%s)\n", error_message, err.Error())
		return DeviceList{}, errors.New(error_message)
	}

	return device_list, nil
}

func GetDeviceList(cfg *configreader.Config) (DeviceList, error) {
	devices, err := getDeviceListFromServer(cfg)
	if err != nil {
		return DeviceList{}, err
	}

	return devices, nil
}

func DescribeDeviceByDevID(cfg *configreader.Config, device_id string) (Device, error) {
	device_list, err := GetDeviceList(cfg)
	if err != nil {
		return Device{}, err
	}

	for _, device := range device_list.Devices {
		if device.DeviceID == device_id {
			return device, nil
		}
	}

	return Device{}, nil
}

func DescribeSystemDeviceByDevID(cfg *configreader.Config, device_id string) (SystemDeviceList, error) {
	device, err := DescribeDeviceByDevID(cfg, device_id)
	if err != nil {
		return SystemDeviceList{}, err
	}

	system_ip := device.SystemIP

	content, err := sendRequestToServer(cfg, "GET", cfg.Server+"/dataservice/system/device/vedges?deviceIP="+system_ip, "application/json", "")
	if err != nil {
		return SystemDeviceList{}, err
	}

	system_device_list := SystemDeviceList{}
	if err := json.Unmarshal(content, &system_device_list); err != nil {
		error_message := "failed to unmarshal system device list"
		fmt.Printf("ERROR: %s (%s)\n", error_message, err.Error())
		return SystemDeviceList{}, errors.New(error_message)
	}

	return system_device_list, nil
}
