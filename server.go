package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	Name 		 string
	URL 		 string
	ReverseProxy *httputil.ReverseProxy
	Health 		 bool
}

func newServer(name string, urlStr string) *server{
	url,_ := url.Parse(urlStr)
	rProxy := httputil.NewSingleHostReverseProxy(url)
	return &server{
		Name : name,
		URL : urlStr,
		ReverseProxy: rProxy,
		Health: true,
	}
}

func (s *server) checkHealth() bool {
	resp, err := http.Head(s.URL)
	if err != nil  ||  resp.StatusCode != http.StatusOK {
		s.Health = false
	}else {
		s.Health = true
	}
	return s.Health
}