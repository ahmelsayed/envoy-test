package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	appVersion := os.Getenv("APP_VERSION")
	if appVersion == "" {
		appVersion = "NO VALUE SET"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port: %s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s 200\n", r.Method, r.URL.String())
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>App '%s' listening on %s</h1>", appVersion, port)
	})
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
