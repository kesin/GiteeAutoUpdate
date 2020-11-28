package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	config   Config
	projects Projects
)

func main() {
	// parse args and init config file
	var configPath string
	args := os.Args
	processArgs(args, &configPath, &config)

	// init projects
	refreshWhitelist()

	// start http server
	http.Handle("/sync", RequestMiddleware(Sync))
	http.Handle("/update_whitelist", RequestMiddleware(UpdateWhitelist))

	addr := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("GAU server listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func RequestMiddleware(f func(http.ResponseWriter, *http.Request)) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %s from %s via %s", r.RequestURI, r.RemoteAddr, r.Header.Get("User-agent"))
		if err := checkIpWhitelist(r.RemoteAddr); err != nil {
			statusCodeWithMessage(&w, 403, err.Error())
			return
		}
		f(w, r)
	}
	return http.HandlerFunc(h)
}

func Sync(w http.ResponseWriter, r *http.Request) {
	// get hook
	hook, err := getHook(r)
	if err != nil {
		statusCodeWithMessage(&w, 500, err.Error())
		return
	}

	// judge and process
	giteeUrl := hook.Repo["url"].(string)
	if isValidRepo(giteeUrl) {
		statusCodeWithMessage(&w, 200, "Received sync request, processing...")
	} else {
		statusCodeWithMessage(&w, 403, "This repo is not in whitelist, "+
			"please create a PullRequest at https://gitee.com/kesin/GiteeAutoUpdate to add your repo")
	}
}

func UpdateWhitelist(w http.ResponseWriter, r *http.Request) {
	// get hook
	hook, err := getHook(r)
	if err != nil {
		statusCodeWithMessage(&w, 500, err.Error())
		return
	}

	// check token
	token := r.FormValue("token")
	if token != config.UpdateWhitelistToken {
		statusCodeWithMessage(&w, 401, "Unauthorized")
		return
	}

	rawWhitelist := fmt.Sprintf("%s/raw/master/syncWhitelist", hook.Repo["url"])
	refreshWhitelist()
	statusCodeWithMessage(&w, 200, rawWhitelist)
}
