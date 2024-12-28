# Load Balancer

## Overview

This project implements a basic TCP-based load balancer in Go. The load balancer distributes incoming client requests across multiple backend servers using configurable strategies like Round Robin and Least Connections. It also includes connection pooling to optimize backend resource usage and supports HTTP request forwarding.

## Features

- **Routing Strategies**: Round Robin and Least Connections.
- **Connection Pooling**: Efficiently reuse backend connections.
- **Error Handling**: Returns appropriate responses when no backends are available.
- **Extensibility**: Easily add more routing strategies or middleware.

## Architecture

The load balancer listens for incoming client TCP connections and forwards the requests to backend servers. It uses:

- **Strategies**: Determines which backend server to forward requests to.
- **Connection Pool**: Manages backend connections to reduce latency and improve throughput.

### Folder Structure

```
.
├── client
│   └── main.go
├── go.mod
├── go.sum
├── interface
│   └── stategy.go
├── main.go
├── models
│   └── server.go
├── package
│   └── minheap.go
├── services
│   ├── connectionpool
│   │   └── pool.go
│   └── loadbalancer.go
└── strategies
    ├── IpHash.go
    ├── LeastConnection.go
    ├── RoundRobin.go
    └── roundRobin_test.go
```

## Requirements

- Go 1.19 or later

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/load-balancer.git
   cd load-balancer
   ```
2. Build the project:
   ```bash
   go build
   ```

## Usage

### Start the Load Balancer

1. Configure the backend servers and strategy in `main.go`:
   ```go
   servers := []string{"127.0.0.1:8081", "127.0.0.1:8082"}
   strategy := strategies.NewRoundRobinStrategy()
   loadBalancer := services.NewLoadBalancer(servers, 8080, strategy)
   loadBalancer.Start()
   ```
2. Run the load balancer:
   ```bash
   ./load-balancer
   ```

### Test the Load Balancer

#### Using `curl`:

```bash
curl --location 'localhost:8080' --header 'Content-Type: application/json' --data '{"key":"value"}'
```

#### Using Apache Benchmark:

```bash
ab -n 1000 -c 100 http://localhost:8080/
```

## Extending the Load Balancer

### Add a New Strategy

1. Create a new file in the `strategies` folder (e.g., `hashing.go`).
2. Implement the `interfaces.Strategy` interface:
   ```go
   type HashingStrategy struct {}

   func (h *HashingStrategy) AddServer(address string) {}
   func (h *HashingStrategy) RemoveServer(address string) {}
   func (h *HashingStrategy) NextServer(conn net.Conn) string {}
   ```
3. Use your strategy when initializing the load balancer.

### Customize Connection Pool

Modify `connection_pool.go` to tweak pool size or implement custom connection re-use logic.

## Future Enhancements

- Support for HTTPS traffic.
- Health checks for backend servers.
- Weighted routing strategies.
- Metrics and logging with Prometheus.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
---

Feel free to reach out with any questions or feedback!