package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	var counter int
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, "Hitting backend %s (count %d)\n", localIP, counter)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
