package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os/exec"
	"strings"
)

type GenConfig struct {
	Dsn    string `yaml:"dsn"`
	Db     string `yaml:"db"`
	Tables string `yaml:"tables"`
}

func getGenConfig() GenConfig {
	data, err := ioutil.ReadFile("./gen.yaml")
	if err != nil {
		panic(fmt.Sprintf("Get gen config error: %s", err.Error()))
	}
	var config GenConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(fmt.Sprintf("Gen config conversion error: %s", err.Error()))
	}
	return config
}

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %s\n", err)
		}
	}()

	genConfig := getGenConfig()

	if genConfig.Tables == "" {
		panic("Please configure Tables")
	}

	for _, table := range strings.Split(genConfig.Tables, ",") {
		tableModel := fmt.Sprintf("%s_model", table)
		outPath := fmt.Sprintf("./internal/model/%s", tableModel)
		cmd := exec.Command("gentool",
			"-dsn", genConfig.Dsn,
			"-tables", table,
			"-db", genConfig.Db,
			"-outPath", outPath,
			"-modelPkgName", tableModel,
			"-onlyModel")
		_, err := cmd.Output()

		if err != nil {
			panic(fmt.Sprintf("Command error: %s", err.Error()))
		}

		fmt.Printf("\033[32m%s\033[0m\n", "Successfully generated table!")
	}
}
