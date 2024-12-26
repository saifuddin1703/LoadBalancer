package strategies

import (
	"container/heap"
	"load-balancer/models"
	pkg "load-balancer/package"
	"sync"
)

type LeastConnectionStrategy struct {
	Servers *pkg.MinHeap
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
	server := models.BackendServer{
		Address: address,
	}
	heap.Push(s.Servers, server)
	// s.Servers = append(s.Servers, address)
}

func (s *LeastConnectionStrategy) RemoveServer(address string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Iterate through the heap to find the server and remove it
	for i, server := range *s.Servers {
		if server.Address == address {
			// Swap with last element and pop
			(*s.Servers)[i], (*s.Servers)[len(*s.Servers)-1] = (*s.Servers)[len(*s.Servers)-1], (*s.Servers)[i]
			(*s.Servers) = (*s.Servers)[:len(*s.Servers)-1]
			heap.Init(s.Servers) // Re-heapify
			break
		}
	}
}

func (s *LeastConnectionStrategy) NextServer() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Get the server with the least connections
	server := heap.Pop(s.Servers).(*models.BackendServer)

	// Return the address of the selected server
	s.UpdateConnectionCount(server.Address, server.ConnectionCount+1)
	return server.Address
}

func (s *LeastConnectionStrategy) UpdateConnectionCount(address string, newCount int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Find the server in the heap and update the connection count
	for _, server := range *s.Servers {
		if server.Address == address {
			s.Servers.UpdateConnectionCount(&server, newCount)
			break
		}
	}
}
