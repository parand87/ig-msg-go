package instagram

import (
	"gopkg.in/yaml.v3"
	"os"
)

var configFileName string

type Config struct {
	AppId           string `yaml:"app_id"`
	AppSecret       string `yaml:"app_secret"`
	RedirectUri     string `yaml:"redirect_uri"`
	GrantType       string `yaml:"grant_type,omitempty"`
	ResponseType    string `yaml:"response_type,omitempty"`
	Domain          string `yaml:"domain,omitempty"`
	FBLoginDomain   string `yaml:"fb_login_domain,omitempty"`
	Prefix          string `yaml:"prefix,omitempty"`
	AccessToken     string `yaml:"access_token,omitempty"`
	PageAccessToken string `yaml:"page_access_token,omitempty"`
}

func setConfigFileName(fileName string) {
	configFileName = fileName
}

func LoadConfig(fileName string) (*Config, error) {
	setConfigFileName(fileName)
	config := &Config{}

	file, err := os.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func SaveConfig(config *Config) error {
	existingConfig, err := LoadConfig(configFileName)
	if err != nil {
		return err
	}

	existingConfig.AccessToken = config.AccessToken

	data, err := yaml.Marshal(existingConfig)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFileName, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
