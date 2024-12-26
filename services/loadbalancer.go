package services

import (
	interfaces "load-balancer/interface"
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
}
func (lb *LoadBalancer) AddServer(address string) {
	lb.strategy.AddServer(address)
}

func (lb *LoadBalancer) RemoveServer(address string) {
	lb.strategy.RemoveServer(address)
}
