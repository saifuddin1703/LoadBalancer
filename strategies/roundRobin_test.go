package strategies

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddServer(t *testing.T) {
	strategy := NewRoundRobinStrategy()
	strategy.AddServer("192.168.0.1:8080")
	strategy.AddServer("192.168.0.2:8080")

	assert.Equal(t, 2, len(strategy.Servers))
	assert.Equal(t, "192.168.0.1:8080", strategy.Servers[0])
	assert.Equal(t, "192.168.0.2:8080", strategy.Servers[1])
}

func TestRemoveServer(t *testing.T) {
	strategy := NewRoundRobinStrategy()
	strategy.AddServer("192.168.0.1:8080")
	strategy.AddServer("192.168.0.2:8080")
	strategy.AddServer("192.168.0.3:8080")

	strategy.RemoveServer("192.168.0.2:8080")

	assert.Equal(t, 2, len(strategy.Servers))
	assert.Equal(t, "192.168.0.1:8080", strategy.Servers[0])
	assert.Equal(t, "192.168.0.3:8080", strategy.Servers[1])
}

func TestNextServer(t *testing.T) {
	strategy := NewRoundRobinStrategy()
	strategy.AddServer("192.168.0.1:8080")
	strategy.AddServer("192.168.0.2:8080")
	strategy.AddServer("192.168.0.3:8080")

	// Mock connection (not used in the method but needed for the signature)

	// Expected order is round-robin: 1 → 2 → 3 → 1 → ...
	assert.Equal(t, "192.168.0.1:8080", strategy.NextServer())
	assert.Equal(t, "192.168.0.2:8080", strategy.NextServer())
	assert.Equal(t, "192.168.0.3:8080", strategy.NextServer())
	assert.Equal(t, "192.168.0.1:8080", strategy.NextServer())
}

func TestNextServerEmptyList(t *testing.T) {
	strategy := NewRoundRobinStrategy()

	// Mock connection

	// With no servers, NextServer should return an empty string
	assert.Equal(t, "", strategy.NextServer())
}

func TestConcurrentAddRemoveNextServer(t *testing.T) {
	strategy := NewRoundRobinStrategy()
	strategy.AddServer("192.168.0.1:8080")
	strategy.AddServer("192.168.0.2:8080")

	done := make(chan bool)

	// Concurrently add servers
	go func() {
		for i := 3; i <= 10; i++ {
			strategy.AddServer("192.168.0." + fmt.Sprint(i) + ":8080")
		}
		done <- true
	}()

	// Concurrently remove servers
	go func() {
		strategy.RemoveServer("192.168.0.2:8080")
		done <- true
	}()

	// Concurrently get servers
	go func() {
		for i := 0; i < 10; i++ {
			strategy.NextServer()
		}
		done <- true
	}()

	// Wait for all goroutines to finish
	<-done
	<-done
	<-done

	// Ensure no race conditions occurred
	t.Log("Concurrent operations completed successfully")
}
