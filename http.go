package main

import (
	"log"
	"net/http"
)

type Hook struct {
	Name string `json:"hook_name"`
	Repo map[string]interface{} `json:"repository"`
}

func RequestMiddleware(f func(http.ResponseWriter, *http.Request)) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %s from %s via %s", r.RequestURI, r.RemoteAddr, r.Header.Get("User-agent") )
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
		statusCodeWithMessage(&w, 403, "This repo is not in whitelist, " +
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

	// judge and process
	giteeUrl := hook.Repo["url"].(string)
	if isValidRepo(giteeUrl) {
		statusCodeWithMessage(&w, 200, "Received sync request, processing...")
	} else {
		statusCodeWithMessage(&w, 403, "Operation not permitted")
	}
}