package oxy

import (
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/roundrobin"
	"net/http"
	"net/url"
	"testing"
)

// 反向代理
func TestName(t *testing.T) {
	fwd, _ := forward.New()
	urlParse, _ := url.Parse("http://localhost:63450")
	redirect := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// let us forward this request to another server
		req.URL = urlParse
		fwd.ServeHTTP(w, req)
	})

	// that's it! our reverse proxy is ready!
	s := &http.Server{
		Addr:    ":8080",
		Handler: redirect,
	}
	s.ListenAndServe()
}

// 负载均衡
func TestOxy(t *testing.T) {

	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	fwd, _ := forward.New()
	lb, _ := roundrobin.New(fwd)

	//lb.UpsertServer(url1)
	//lb.UpsertServer(url2)

	s := &http.Server{
		Addr:    ":8080",
		Handler: lb,
	}
	s.ListenAndServe()
}
