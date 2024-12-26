package interfaces

import "net"

type Strategy interface {
	AddServer(address string)
	RemoveServer(address string)
	NextServer(c net.Conn) string
}
