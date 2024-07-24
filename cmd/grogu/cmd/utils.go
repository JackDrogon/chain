package cmd

import (
	"fmt"

	"github.com/cometbft/cometbft/rpc/client/http"
)

func createClients(nodeURIs []string) ([]*http.HTTP, func(), error) {
	clients := make([]*http.HTTP, 0, len(nodeURIs))
	for _, uri := range nodeURIs {
		httpClient, err := http.New(uri, "/websocket")
		if err != nil {
			fmt.Printf("[Grogu] failed to create HTTP client with error: %v\n", err)
			continue
		}

		if err = httpClient.Start(); err != nil {
			fmt.Printf("[Grogu] failed to start HTTP client with error: %v\n", err)
			continue
		}

		clients = append(clients, httpClient)
	}

	if len(clients) == 0 {
		return nil, nil, fmt.Errorf("no clients are available")
	}

	// Function to stop all clients created so far
	stopClients := func() {
		for _, client := range clients {
			_ = client.Stop()
		}
	}

	return clients, stopClients, nil
}
