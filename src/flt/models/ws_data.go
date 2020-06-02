/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package models

import (
	"github.com/gorilla/websocket"
	"sync"
	"fmt"
)

var once sync.Once

type Message struct {
    Message string `json:"message"`
    Time string
}

// type global
type singleton struct {
	Clients map[*websocket.Conn]bool
	Broadcast chan Message
}

var (
	instance singleton
)

func WsData() singleton {

	once.Do(func() { // <-- atomic, does not allow repeating
		instance = singleton{ Clients: make(map[*websocket.Conn]bool), Broadcast: make(chan Message) } // <-- thread safe
		//go SendBroadcasts()
	})

	return instance
}

func SendBroadcastMessage(msg Message) {
	var data = WsData()
	for client := range data.Clients {
		err := client.WriteJSON(msg)
		if err != nil {
			fmt.Println("client.WriteJSON error: %v", err)
			client.Close()
			delete(data.Clients, client)
		}
	}
}

func SendBroadcasts() {
	var data = WsData()
	for {
		msg := <-data.Broadcast
		fmt.Println("clients len ", len(data.Clients))
		SendBroadcastMessage(msg)
	}
}
