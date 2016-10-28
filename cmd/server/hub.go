// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)
// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {

	// Registered clients.
	clients map[*Client]bool

}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
	}
}

// Register a client into the Hub
func (h *Hub) RegisterClient(client *Client) {
	log.Printf("Registering client from %s", client.Ip)
	h.clients[client] = true
}

// Unregister a client from the hub
func (h *Hub) UnregisterClient(client *Client) {
	if _, ok := h.clients[client]; ok {
		delete(h.clients, client)
	}
}

// Main loop of the Hub. This will listen for events on the input channel, and then broadcast all
// vents/message to all listenning clients
func (h *Hub) run(ic chan DataPoint) {
	for {
		select {
		case p := <-ic:
			h.SendValue(p)
		}
	}
}

// Send a value to all registered clients
func (h *Hub) SendValue(p DataPoint) {
	for cl := range h.clients {
		err := cl.send(p)
		if err != nil {
			log.Println("Error sending to client. Maybe client got away")
			h.UnregisterClient(cl)
		}
	}
}

// Return a slice of client.
func (h *Hub) GetClients() []*Client {
	clients := make([]*Client, len(h.clients))

	i := 0
	for cl := range h.clients {
		clients[i] = cl
		i++;
	}

	return clients
}


