package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	sproxyurl := flag.String("url", "https://usas4021.phx-dc.dhl.com:12422", "target URL")
	flag.Parse()
	// Parse the target URL
	target, err := url.Parse(*sproxyurl)
	if err != nil {
		panic(err)
	}

	// Create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Start the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Update the request URL to match the target URL
		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = target.Host

		// Forward the request to the target server
		proxy.ServeHTTP(w, r)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
