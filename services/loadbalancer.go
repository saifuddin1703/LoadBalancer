package services

import (
	"fmt"
	"io"
	interfaces "load-balancer/interface"
	"load-balancer/services/connectionpool"
	"net"
	"sync"
	"time"

	"github.com/labstack/gommon/log"
)

type LoadBalancer struct {
	Servers  []string
	pools    map[string]*connectionpool.ConnectionPool
	poolSize int
	Port     int
	mu       sync.Mutex
	strategy interfaces.Strategy
}

func NewLoadBalancer(servers []string, port int, strategy interfaces.Strategy) *LoadBalancer {
	lb := LoadBalancer{
		Servers:  servers,
		Port:     port,
		strategy: strategy,
		mu:       sync.Mutex{},
		pools:    make(map[string]*connectionpool.ConnectionPool),
		poolSize: 500,
	}

	for _, server := range servers {
		lb.AddServer(server)
	}
	return &lb
}

func (lb *LoadBalancer) Start() {
	// start tcp server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", lb.Port))
	if err != nil {
		log.Fatalf("Failed to start load balancer: %v", err)
	}
	defer listener.Close()

	log.Printf("Load balancer started on port %d", lb.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go lb.forwarRequest(conn)
	}
}
func (lb *LoadBalancer) forwarRequest(conn net.Conn) {
	backendAddr := lb.strategy.NextServer()
	if backendAddr == "" {
		log.Info("No available backend servers")
		conn.Write([]byte("503 Service Unavailable"))
		return
	}

	backendPool := lb.GetOrCreatePool(backendAddr)
	backendConn, err := backendPool.Acquire()
	backendConn.SetDeadline(time.Now().Add(5 * time.Minute))
	if err != nil {
		log.Printf("Failed to connect to backend server %s: %v", backendAddr, err)
		conn.Write([]byte("503 Service Unavailable"))
		return
	}
	defer backendPool.Release(backendConn) // closing the backend's connection
	defer conn.Close()                     // closing the client's connection

	// Relay data between client and backend
	go io.Copy(backendConn, conn)
	io.Copy(conn, backendConn)
}

func (lb *LoadBalancer) GetOrCreatePool(address string) *connectionpool.ConnectionPool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if _, exists := lb.pools[address]; !exists {
		lb.pools[address] = connectionpool.NewConnectionPool(address, lb.poolSize)
	}
	return lb.pools[address]
}
func (lb *LoadBalancer) AddServer(address string) {
	lb.strategy.AddServer(address)
}

func (lb *LoadBalancer) RemoveServer(address string) {
	lb.strategy.RemoveServer(address)
}
