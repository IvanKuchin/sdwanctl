package config_reader

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func parseConfigFile(config_file Config_file) (Config, error) {
	var cfg Config
	var user_name string
	var vmanage_name string

	found := false
	for _, context := range config_file.Contexts {
		if context.Name == config_file.Current_context {
			user_name = context.User
			vmanage_name = context.Vmanage
			found = true
			break
		}
	}

	if !found {
		error_message := "ERROR: No context found in the config file"
		log.Printf("ERROR: %s", error_message)
		return cfg, errors.New(error_message)
	}

	found = false
	for _, vmanage := range config_file.Vmanages {
		if vmanage.Name == vmanage_name {
			cfg.Server = vmanage.Server
			cfg.BearerToken = vmanage.BearerToken
			found = true
		}
	}

	if !found {
		error_message := "ERROR: No vmanage found in the config file"
		log.Printf("ERROR: %s", error_message)
		return cfg, errors.New(error_message)
	}

	found = false
	for _, user := range config_file.Users {
		if user.Name == user_name {
			cfg.Login = user.Login
			cfg.Password = user.Password
			found = true
		}
	}

	if !found {
		error_message := "ERROR: No user found in the config file"
		log.Printf("ERROR: %s", error_message)
		return cfg, errors.New(error_message)
	}

	return cfg, nil
}

type YAMLConfigReader struct {
}

func (ycr *YAMLConfigReader) GetConfig() (*Config, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("ERROR: %v\n", err.Error())
	}

	filename := filepath.Join(home_dir, ".sdwan", "config")

	if _, err := os.Stat(filename); err != nil {
		log.Fatalf("ERROR: %v\n", err.Error())
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err.Error())
	}

	config_file := Config_file{}
	if err := yaml.Unmarshal(content, &config_file); err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}

	cfg, err := parseConfigFile(config_file)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func NewYAMLConfigReader() (ConfigReader, error) {
	return &YAMLConfigReader{}, nil
}
