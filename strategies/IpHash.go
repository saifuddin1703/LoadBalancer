package strategies

import "net"

type IpHashStrategy struct {
	Servers []string
	Port    int
}

func (s *IpHashStrategy) AddServer(address string) {
	
}
func (s *IpHashStrategy) RemoveServer(address string)
func (s *IpHashStrategy) NextServer(c net.Conn) string
