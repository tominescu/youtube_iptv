package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	log "github.com/tominescu/double-golang/simplelog"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(503), 503)
		return
	}
	id := r.Form.Get("id")
	if id == "" {
		http.Error(w, http.StatusText(503), 503)
		return
	}
	q := r.Form.Get("q")
	url := "https://www.youtube.com/watch?v=" + id
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	re := regexp.MustCompile(`hlsvp.*m3u8`)
	hls := re.Find(body)
	if len(hls) < 8 {
		http.Error(w, "Cant't find m3u8 url", 503)
		return
	}
	dst := string(hls[8:])
	dst = strings.Replace(dst, "\\/", "/", -1)
	if q == "" {
		dst = strings.Replace(dst, "https://manifest.googlevideo.com", "http://"+r.Host, 1)
		w.Header().Set("Location", dst)
		http.Error(w, http.StatusText(302), 302)
		return
	}
	seq, err := strconv.Atoi(q)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	resp2, err := http.Get(dst)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp2.Body.Close()
	body, err = ioutil.ReadAll(resp2.Body)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	re = regexp.MustCompile(`https://.*\.m3u8`)
	urls := re.FindAllString(string(body), -1)
	seq = len(urls) - seq
	if seq < 0 || seq >= len(urls) {
		errStr := fmt.Sprintf("Quality not exist, total: %d", len(urls))
		http.Error(w, errStr, 503)
		return
	}

	dst = strings.Replace(urls[seq], "https://manifest.googlevideo.com", "http://"+r.Host, 1)
	w.Header().Set("Location", dst)
	http.Error(w, http.StatusText(302), 302)
}
