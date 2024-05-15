package slbKit

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

// Backend holds the data about a server
type Backend struct {
	mux sync.RWMutex

	URL          *url.URL
	Alive        bool
	ReverseProxy *httputil.ReverseProxy
}

// SetAlive for this backend
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// IsAlive returns true when backend is alive
func (b *Backend) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return
}
