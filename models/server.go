package models

// BackendServer represents a backend server with a connection count
type BackendServer struct {
	Address         string
	ConnectionCount int
	Index           int // Index in the heap
}
