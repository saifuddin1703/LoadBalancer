package pkg

import (
	"container/heap"
)

// BackendServer represents a backend server with a connection count
type BackendServer struct {
	Address         string
	ConnectionCount int
	Index           int // Index in the heap
}

// MinHeap is a heap-based priority queue that stores BackendServer objects
type MinHeap []BackendServer

// Implement heap.Interface for MinHeap

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	// The server with fewer connections has a higher priority (min-heap)
	return h[i].ConnectionCount < h[j].ConnectionCount
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

// Push adds a BackendServer to the heap
func (h *MinHeap) Push(x interface{}) {
	n := len(*h)
	server := x.(*BackendServer)
	server.Index = n
	*h = append(*h, *server)
}

// Pop removes and returns the smallest BackendServer (with least connections)
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	server := old[n-1]
	*h = old[0 : n-1]
	return &server
}

// UpdateConnectionCount updates the connection count for a server and re-heapifies
func (h *MinHeap) UpdateConnectionCount(server *BackendServer, newCount int) {
	server.ConnectionCount = newCount
	heap.Fix(h, server.Index)
}
