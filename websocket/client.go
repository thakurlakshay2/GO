package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct{
	connection *websocket.Conn
	manager *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client{
	return &Client{
		connection:conn,
		manager: manager,
	}
}

func (c *Client) readMessages(){
	defer func(){
		//CLEAN UP CONNECTION
		c.manager.removeClient(c)
	}()
	for{
		messageType,payload,err :=c.connection.ReadMessage()
		if err!=nil {
			if websocket.IsUnexpectedCloseError(err,websocket.CloseGoingAway,websocket.CloseAbnormalClosure){
				log.Printf("error reading message : %v ", err)
			}
			break
		}

		log.Println(messageType)
		log.Println(string(payload))
	}
}