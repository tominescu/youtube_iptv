package main

import (
	"net/http"

	log "github.com/tominescu/double-golang/simplelog"
)

func main() {
	log.SetLevel(log.DEBUG)
	http.HandleFunc("/index.m3u8", indexHandler)
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/videoplayback/", videoHandler)
	http.ListenAndServe(":8080", nil)
}
