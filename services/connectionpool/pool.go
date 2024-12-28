package connectionpool

import (
	"net"
	"sync"
)

type ConnectionPool struct {
	Address     string
	Connections chan net.Conn
	MaxPoolSize int
	mu          sync.Mutex
}

func NewConnectionPool(address string, maxPoolSize int) *ConnectionPool {
	return &ConnectionPool{
		Address:     address,
		Connections: make(chan net.Conn, maxPoolSize),
		MaxPoolSize: maxPoolSize,
	}
}

func (p *ConnectionPool) Acquire() (net.Conn, error) {
	select {
	case conn := <-p.Connections:
		return conn, nil // Reuse an existing connection
	default:
		return net.Dial("tcp", p.Address) // Create a new connection
	}
}

func (p *ConnectionPool) Release(conn net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	select {
	case p.Connections <- conn:
		// Connection returned to pool
	default:
		// Pool is full, close connection
		conn.Close()
	}
}
