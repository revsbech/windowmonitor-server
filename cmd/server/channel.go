package main

import (
	"log"
)

// This represent a channel in the server
type Channel struct {
	id   string
	name string
	hub  *Hub
}

func NewChannel(id string, name string) *Channel {
	var channel = &Channel{
		id:    id,
		name:  name,
		hub:   newHub(),
	}
	log.Printf("New channel %s\n", channel.name)
	c := make(chan DataPoint, 100)
	go channel.hub.run(c)

	// Now, add dummy events on this channel
	go generateEvents(c)
	return channel
}

func (c *Channel) RegisterClient(client *Client) {
	c.hub.RegisterClient(client)
}

func (c *Channel) GetClients() []*Client {
	return c.hub.GetClients()
}


type ChannelView struct {
	name     string         `json:"ip"`
    clients  []*ClientView  `json:"clients"`
}


