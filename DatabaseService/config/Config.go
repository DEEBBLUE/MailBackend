package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct{
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	DBName string `yaml:"dbname"`
}

func LoadConfig(path string) (DatabaseConfig,error){
	var conf DatabaseConfig
	data,err := ioutil.ReadFile(path)

	if err != nil {
		return conf,fmt.Errorf("Config file is empty %w",err)
	}
	
	if err :=	yaml.Unmarshal(data, &conf); err != nil{
		return conf,fmt.Errorf("Config not typed %w",err)
	}
	
	return conf,nil
}

func(conf *DatabaseConfig) GetStringConf() string{
	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",conf.Host,conf.Port,conf.User,conf.Password,conf.DBName) 
	fmt.Println(str)
	fmt.Println(conf)

	return str
}
