package main

import (
	"load-balancer/services"
	"load-balancer/strategies"

	"github.com/labstack/gommon/log"
)

func main() {
	log.Info("Welcome to the Load Balancer")
	address := []string{"localhost:6000", "localhost:6001", "localhost:6002"}
	_ = address
	lb := services.NewLoadBalancer(address, 8080, strategies.NewRoundRobinStrategy())
	lb.Start()
}

// func handleRead(c net.Conn) {
// 	// defer c.Close()

// 	// Use bufio.Reader to read the request line-by-line
// 	reader := bufio.NewReader(c)

// 	// Read the request line (e.g., "GET / HTTP/1.1")
// 	requestLine, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error reading request line:", err)
// 		return
// 	}
// 	fmt.Println("Request Line:", requestLine)

// 	// fmt.Println("string : ", reader.ReadString())
// 	// Read headers until the blank line
// 	headers := make(map[string]string)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		fmt.Println("line : ", line)
// 		if err != nil {
// 			fmt.Println("Error reading headers:", err)
// 			return
// 		}
// 		line = strings.TrimSpace(line)
// 		if line == "" {
// 			break // End of headers
// 		}

// 		// Split header into key-value pair
// 		parts := strings.SplitN(line, ":", 2)
// 		if len(parts) == 2 {
// 			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
// 		}
// 	}
// 	fmt.Println("Headers:", headers)

// 	// If Content-Length header is present, read the body
// 	contentLength := headers["Content-Length"]
// 	if contentLength != "" {
// 		length := 0
// 		fmt.Sscanf(contentLength, "%d", &length)

// 		// Read the exact number of bytes specified by Content-Length
// 		body := make([]byte, length)
// 		_, err := reader.Read(body)
// 		if err != nil {
// 			fmt.Println("Error reading body:", err)
// 			return
// 		}
// 		fmt.Println("Request Body:", string(body))
// 	} else {
// 		fmt.Println("No request body found")
// 	}
// }
