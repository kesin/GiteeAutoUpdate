package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const gauVERSION = "1.0.0"

type Config struct {
	Port              int                 `json:"Port"`
	GiteePrivateToken string              `json:"GiteePrivateToken"`
	IpWhitelist       []string            `json:"IpWhitelist"`
	SyncUser          map[string]SyncUser `json:"SyncUser"`
}

type SyncUser struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func processArgs(args []string, configPath *string, config *Config) {
	// no valid args
	if len(args) < 2 {
		usage()
	}

	// proceed by args
	switch args[1] {
	case "-h":
		usage()
	case "-c":
		getConfig(args, configPath, config)
	case "-v":
		version()
	default:
		usage()
	}
}

func usage() {
	fmt.Printf(`GAU %s - GiteeAutoUpdate is a bot to do something useful for project on Gitee

Usage: Config your config.json and start GAU with command ./gau -c ./config.json &

Args:
	-c	Specify a config json
	-h	Help and quit
	-v	Show current version and quit

`, gauVERSION)
	os.Exit(0)
}

func version() {
	fmt.Printf(`GAU %s - Gitee Auto Update is a bot to do something useful for project on Gitee
Version: %s
`, gauVERSION)
	os.Exit(0)
}

func getConfig(args []string, configPath *string, config *Config) {
	if len(args) < 3 {
		fmt.Println("Please specify config file path.")
		os.Exit(0)
	}

	// check config path input
	configFilePath := args[2]
	cFile, err := os.Stat(configFilePath)
	if err != nil {
		fmt.Printf("Config file is invalid: %s \n", err.Error())
		os.Exit(0)
	}

	// judge if file is a regular file
	if cFile.Mode().IsRegular() {
		*configPath = configFilePath
	} else {
		fmt.Println("Config file is invalid, maybe the path you specified is a directory?")
		os.Exit(0)
	}

	parseConfig(configPath, config)
}

func parseConfig(configPath *string, config *Config) {
	// get File
	configFile, err := os.Open(*configPath)
	if err != nil {
		fmt.Printf("Config file failed to open: %s \n", err.Error())
		os.Exit(0)
	}
	defer func() {
		if err := configFile.Close(); err != nil {
			fmt.Printf("Config file failed to close: %s \n", err.Error())
			os.Exit(0)
		}
	}()
	fmt.Printf("Using config file: %s \n", *configPath)

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(config); err != nil {
		fmt.Printf("Failed to parse config: %s \n", err.Error())
		os.Exit(0)
	}
}