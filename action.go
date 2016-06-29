package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func doAction(c *cli.Context) error {
	fmt.Printf("use %v as config file\n", c.String("config"))

	configPath := c.String("config")
	config, err := readConfig(configPath)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	packageName := c.String("name")
	if packageName == "" {
		log.Fatalln("package name is not provided")
	}
	config.PackageName = packageName

	storeURL := getPlayStoreURL(packageName)
	log.Printf("fetch app info from %v", storeURL)

	config.PlayStoreURL = storeURL

	report("test", config)

	return nil
}

// read config file from provided path
func readConfig(path string) (Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	output := Config{}

	if err != nil {
		return output, err
	}

	err = yaml.Unmarshal(yamlFile, &output)
	if err != nil {
		return output, err
	}

	return output, nil
}

func getPlayStoreURL(packageName string) string {
	return "https://play.google.com/store/apps/details?id=" + packageName
}
