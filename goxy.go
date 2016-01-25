package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

var (
	port = flag.Int("port", 8080, "Port number to start proxy on")
)

func init() {
	flag.Parse()
}

func newDirector(r *http.Request) func(*http.Request) {
	return func(req *http.Request) {
		schemeOveride := r.Header.Get("goxy-scheme-override")

		if schemeOveride != "" {
			req.URL.Scheme = schemeOveride
		} else {
			req.URL.Scheme = "http"
		}

		req.URL.Host = r.Host

		reqLog, err := httputil.DumpRequestOut(req, false)
		if err != nil {
			log.Printf("Got error %s\n %+v\n", err.Error(), req)
		}

		log.Println(string(reqLog))
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	proxy := &httputil.ReverseProxy{
		Transport: &http.Transport{},
		Director:  newDirector(r),
	}
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", proxyHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
