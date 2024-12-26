package strategies

import (
	"sync"
)

type RoundRobinStrategy struct {
	Servers    []string
	Port       int
	LastServer int
	mu         sync.Mutex
}

func NewRoundRobinStrategy() *RoundRobinStrategy {
	return &RoundRobinStrategy{
		LastServer: -1,
		mu:         sync.Mutex{},
	}
}

func (s *RoundRobinStrategy) AddServer(address string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Servers = append(s.Servers, address)
}

func (s *RoundRobinStrategy) RemoveServer(address string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for idx, addr := range s.Servers {
		if addr == address {
			s.Servers = append(s.Servers[:idx], s.Servers[idx+1:]...)
		}
	}
}

func (s *RoundRobinStrategy) NextServer() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.Servers) == 0 {
		return ""
	}
	s.LastServer = (s.LastServer + 1) % len(s.Servers)
	return s.Servers[s.LastServer]
}
