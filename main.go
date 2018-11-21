package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/akamensky/argparse"
)

func handlePreflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func handlePassThrough(w http.ResponseWriter, r *http.Request) {
	target := "http://" + r.URL.Path[1:]
	url, err := url.Parse(target)

	if err != nil {
		log.Println("Could not handle request", *r)
		fmt.Fprintf(w, "Could not parse URL: %s", err)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	r.URL = url
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Host = url.Host
	proxy.ServeHTTP(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request: ", r.Method, r.URL)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		handlePreflight(w, r)
	} else {
		handlePassThrough(w, r)
	}
}

func main() {

	parser := argparse.NewParser("corsproxy", "Add CORS support via a proxy to servers that don't support it.")
	port := parser.String("p", "port", &argparse.Options{Required: false, Help: "The port to listen on", Default: "8080"})
	inface := parser.String("i", "interface", &argparse.Options{Required: false, Help: "The interface to listen on. localhost or 127.0.0.1 or 0.0.0.0 is common. Non routable LAN interface (192.168.X.Y or 10.0.X.Y) is useful. Wan/Public is not recommended.", Default: "127.0.0.1"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	http.HandleFunc("/", handler)

	localServerURL := *inface + ":" + *port
	log.Println("Listening on: ", localServerURL)
	log.Fatal(http.ListenAndServe(localServerURL, nil))
}
