package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	serverList      []*server
	lastServedIndex = 0
)

func main() {
	serverList = []*server{
		newServer("server1", "http://127.0.0.1:5000"),
		newServer("server2", "http://127.0.0.1:5001"),
		newServer("server3", "http://127.0.0.1:5002"),
		newServer("server4", "http://127.0.0.1:5003"),
		newServer("server5", "http://127.0.0.1:5004"),
	}
	http.HandleFunc("/", forwardRequest)
	go startHealthCheck()
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	server, err := getHealthyServer()
	if err != nil {
		http.Error(res, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	server.ReverseProxy.ServeHTTP(res, req)
}

func getHealthyServer() (*server, error) {
	for i := 0; i < len(serverList); i++ {
		server := getServer()
		if server.Health {
			return server, nil
		}
	}
	return nil, fmt.Errorf("no healthy hosts are available")
}

func getServer() *server {
	nextIndex := (lastServedIndex + 1) % len(serverList)
	server := serverList[lastServedIndex]
	lastServedIndex = nextIndex
	return server
}
