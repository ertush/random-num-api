package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func randomNumHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {

		if req.RequestURI == "/random-num" {
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Add("Host", "localhost:8081")

			w.WriteHeader(http.StatusOK)

			min := 20
			max := 10000

			fmt.Fprintf(w, "%v", rand.Intn(max-min)+min)
		}

	} else {
		http.NotFound(w, req)
	}

}

// GetPort is not exported
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func main() {

	http.HandleFunc("/random-num", randomNumHandler)

	s := &http.Server{
		Addr:           GetPort(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 Mib
	}

	log.Println("[+] Server is listening on localhost:8081")
	log.Fatal(s.ListenAndServe())

}
