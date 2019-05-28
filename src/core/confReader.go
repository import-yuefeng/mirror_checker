package core

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (imageList *Images) GetConf(confPath string) *Images {
	// yamlFile, err := ioutil.ReadFile("/etc/mirror_checker/default.yaml")
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, imageList)
	if err != nil {
		fmt.Println(err.Error())
	}
	return imageList
}
