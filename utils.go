package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkIpWhitelist(refer string) error {
	referURL := strings.Split(refer, ":")
	ip := referURL[0]
	if inWhitelist(config.IpWhitelist, ip) {
		return nil
	} else {
		return errors.New("Request ip not in whitelist\n")
	}
}

func statusCodeWithMessage(w *http.ResponseWriter, code int, message string) {
	(*w).WriteHeader(code)
	message = fmt.Sprintf("%s\n", message)
	_, _ = (*w).Write([]byte(message))
}

func inWhitelist(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getHook(r *http.Request) (hook Hook, err error) {
	jsonParser := json.NewDecoder(r.Body)
	if err := jsonParser.Decode(&hook); err != nil {
		log.Printf("Parse error: %s\n", err.Error())
	}
	return hook, err
}

func isValidRepo(url string) bool {
	for k, _ := range projects {
		if k == url {
			return true
		}
	}
	return false
}

func refreshWhitelist() {
	projects = nil
	jsonFile, err := os.Open("./syncWhitelist")
	if err != nil {
		fmt.Printf("Whitelist file failed to open: %s \n", err.Error())
		os.Exit(0)
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			fmt.Printf("Whitelist file failed to close: %s \n", err.Error())
			os.Exit(0)
		}
	}()

	jsonParser := json.NewDecoder(jsonFile)
	if err := jsonParser.Decode(&projects); err != nil {
		fmt.Printf("Failed to parse Whitelist: %s \n", err.Error())
		os.Exit(0)
	}
}
