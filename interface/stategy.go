package interfaces

type Strategy interface {
	AddServer(address string)
	RemoveServer(address string)
	NextServer() string
}
