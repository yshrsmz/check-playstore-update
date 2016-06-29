package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/roylee0704/gron"
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

	appConfig := createAppConfig(config, c)

	checkInfo := CheckInfo{}

	var wg sync.WaitGroup
	wg.Add(1)
	g := gron.New()
	g.AddFunc(gron.Every(3*time.Second), func() {
		log.Println("checking update...")
		checkInfo, err = check(appConfig.PlayStoreURL, checkInfo)
		if err != nil {
			log.Fatalf("error while checking update: %v", err)
			reportError(appConfig)
			g.Stop()
			wg.Done()
		}

		if checkInfo.IsUpdated {
			reportSuccess(appConfig)
			g.Stop()
			wg.Done()
		}
	})

	g.Start()

	reportStart(appConfig)
	wg.Wait()

	return nil
}

// read config file from provided path
func readConfig(path string) (Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	output := Config{
		SleepTime: 60,
	}

	if err != nil {
		return output, err
	}

	err = yaml.Unmarshal(yamlFile, &output)
	if err != nil {
		return output, err
	}

	return output, nil
}

func createAppConfig(config Config, c *cli.Context) AppConfig {
	packageName := c.String("name")
	if packageName == "" {
		log.Fatalln("package name is not provided")
	}

	storeURL := getPlayStoreURL(packageName)
	log.Printf("fetch app info from %v", storeURL)

	appConfig := AppConfig{
		Params:       config,
		PlayStoreURL: storeURL,
		PackageName:  packageName,
	}

	return appConfig
}

func getPlayStoreURL(packageName string) string {
	return "https://play.google.com/store/apps/details?id=" + packageName
	// return "http://127.0.0.1:8000/test.html"
}
