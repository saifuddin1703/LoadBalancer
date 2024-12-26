package strategies

import (
	"sync"
)

type LeastConnectionStrategy struct {
	Servers []string
	mu      sync.Mutex
}

func NewLeastConnectionStrategy() *LeastConnectionStrategy {
	return &LeastConnectionStrategy{
		mu: sync.Mutex{},
	}
}

func (s *LeastConnectionStrategy) AddServer(address string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Servers = append(s.Servers, address)
}

func (s *LeastConnectionStrategy) RemoveServer(address string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for idx, addr := range s.Servers {
		if addr == address {
			s.Servers = append(s.Servers[:idx], s.Servers[idx+1:]...)
		}
	}
}

func (s *LeastConnectionStrategy) NextServer() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.Servers) == 0 {
		return ""
	}

	// TODO()
	return ""
}
