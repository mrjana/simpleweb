package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	var counter int
	var version = "1.0"
	http.HandleFunc("/count/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := net.Dial("udp", r.RemoteAddr)
		if err != nil {
			log.Fatalf("could not determine local IP address: %v", err)
		}
		localAddr := conn.LocalAddr().String()
		conn.Close()

		localIP, _, err := net.SplitHostPort(localAddr)
		if err != nil {
			log.Fatalf("could not split local IP address: %v", err)
		}

		counter++
		fmt.Fprintf(w, "Hitting backend %s (count %d, version %s)\n", localIP, counter, version)
	})

	http.HandleFunc("/health/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK\n")
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
