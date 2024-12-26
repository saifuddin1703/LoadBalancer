package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/labstack/gommon/log"
)

func main() {
	log.Info("Welcome to the Load Balancer")

	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("Error listening : ", err)
	}
	log.Info("Listening on port 8080")
	for {
		con, _ := conn.Accept()
		go func(c net.Conn) {
			log.Info("Received connection")
			handleRead(c)
			message := "Hellow from load balancer"
			response := fmt.Sprintf(
				"HTTP/1.1 200 OK\r\n"+
					"Content-Type: text/plain\r\n"+
					"Content-Length: %d\r\n"+
					"\r\n"+
					"%s",
				len(message), message,
			)
			// net.Dial()
			fmt.Println("response : ", response)
			c.Write([]byte(response))
			c.Close()
		}(con)
	}
}

func handleRead(c net.Conn) {
	// defer c.Close()

	// Use bufio.Reader to read the request line-by-line
	reader := bufio.NewReader(c)

	// Read the request line (e.g., "GET / HTTP/1.1")
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request line:", err)
		return
	}
	fmt.Println("Request Line:", requestLine)

	// fmt.Println("string : ", reader.ReadString())
	// Read headers until the blank line
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		fmt.Println("line : ", line)
		if err != nil {
			fmt.Println("Error reading headers:", err)
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break // End of headers
		}

		// Split header into key-value pair
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	fmt.Println("Headers:", headers)

	// If Content-Length header is present, read the body
	contentLength := headers["Content-Length"]
	if contentLength != "" {
		length := 0
		fmt.Sscanf(contentLength, "%d", &length)

		// Read the exact number of bytes specified by Content-Length
		body := make([]byte, length)
		_, err := reader.Read(body)
		if err != nil {
			fmt.Println("Error reading body:", err)
			return
		}
		fmt.Println("Request Body:", string(body))
	} else {
		fmt.Println("No request body found")
	}
}
