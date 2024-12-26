package services

import (
	"fmt"
	"io"
	interfaces "load-balancer/interface"
	"net"

	"github.com/labstack/gommon/log"
)

type LoadBalancer struct {
	Servers  []string
	Port     int
	strategy interfaces.Strategy
}

func NewLoadBalancer(servers []string, port int, strategy interfaces.Strategy) *LoadBalancer {
	lb := LoadBalancer{
		Servers:  servers,
		Port:     port,
		strategy: strategy,
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
		go lb.handleConnection(conn)
	}
}
func (lb *LoadBalancer) handleConnection(conn net.Conn) {
	defer conn.Close()
	backendAddr := lb.strategy.NextServer()
	if backendAddr == "" {
		log.Info("No available backend servers")
		conn.Write([]byte("503 Service Unavailable"))
		return
	}

	backendConn, err := net.Dial("tcp", backendAddr)
	if err != nil {
		log.Printf("Failed to connect to backend server %s: %v", backendAddr, err)
		conn.Write([]byte("503 Service Unavailable"))
		return
	}
	defer backendConn.Close()

	// Relay data between client and backend
	go io.Copy(backendConn, conn)
	io.Copy(conn, backendConn)
}
func (lb *LoadBalancer) AddServer(address string) {
	lb.strategy.AddServer(address)
}

func (lb *LoadBalancer) RemoveServer(address string) {
	lb.strategy.RemoveServer(address)
}
