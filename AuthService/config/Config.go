package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)


type AuthConfig struct{
	Credential []byte `yaml:"cred"`
}

func InitConfig(path string)(*AuthConfig,error){
	var res AuthConfig
	data,err := ioutil.ReadFile(path)
	if err != nil {
		return nil,err
	}
	if err := yaml.Unmarshal(data,&res);err != nil {
		return nil,err
	}
	return &res,nil
}
