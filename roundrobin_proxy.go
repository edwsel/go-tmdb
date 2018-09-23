package tmdb

import (
	"sync"
)

type roundRobinProxy struct {
	proxies       []Proxy
	currentTicker int
	maxAllowed    int
	mu            sync.Mutex
}

func InitRoundRobin(proxies []Proxy) (roundRobinProxy) {
	return roundRobinProxy{
		proxies:       proxies,
		maxAllowed:    len(proxies),
		currentTicker: 0,
		mu:            sync.Mutex{},
	}
}

func (r *roundRobinProxy) GetProxy() Proxy {
	for {
		tickerNumber := r.GetTicker()

		proxy := r.proxies[tickerNumber]

		select {
		case <-proxy.throttle:
			return proxy
		default:
			break
		}
	}
}

func (r *roundRobinProxy) GetTicker() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	ticker := r.currentTicker

	if r.currentTicker < r.maxAllowed {
		r.currentTicker = r.currentTicker + 1
	} else {
		r.currentTicker = 0
	}

	return ticker
}

func (r *roundRobinProxy) Next() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.currentTicker < r.maxAllowed {
		r.currentTicker = r.currentTicker + 1
	} else {
		r.currentTicker = 0
	}
}
