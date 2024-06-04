package main

import (
	"github.com/vulcand/oxy/v2/buffer"
	"github.com/vulcand/oxy/v2/forward"
	"github.com/vulcand/oxy/v2/roundrobin"
	"net/http"
	"net/url"
)

func main() {
	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	fwd := forward.New(false)
	lb, err := roundrobin.New(fwd)
	if err != nil {
		panic(err)
	}

	servers := []string{
		"http://localhost:8001",
		"http://localhost:8002",
	}
	for _, s := range servers {
		u, err := url.Parse(s)
		if err != nil {
			panic(err)
		}
		if err := lb.UpsertServer(u); err != nil {
			panic(err)
		}
	}

	// buf will read the request body and will replay the request again in case if forward returned status
	// corresponding to nework error (e.g. Gateway Timeout)
	buf, err := buffer.New(lb, buffer.Retry(`IsNetworkError() && Attempts() < 2`))
	if err != nil {
		panic(err)
	}

	s := &http.Server{
		Addr:    ":8000",
		Handler: buf,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
