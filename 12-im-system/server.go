package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip 			string
	Port 		int

	// storage online user info
	OnlineMap 	map [string] *User
	mapLock		sync.RWMutex

	// publish common message
	PubChan		chan string
}

// Create a server instance
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip: 		ip,
		Port: 		port,
		OnlineMap: 	make(map [string] *User),
		PubChan: 	make(chan string),	
	}

	return server
}

// Listen publish channel
func (this *Server) ListenPublishChannel() {
	for {
		msg := <- this.PubChan

		this.mapLock.Lock()
		for _, user := range this.OnlineMap {
			user.Channel <- msg
		}
		this.mapLock.Unlock()
	}
}

// Broadcast user online information
func (this *Server) Broadcast(user *User, msg string) {
	pubMsg := "[" + user.Addr + "]" + user.Name + msg
	this.PubChan <- pubMsg
}

// Handle connection
func (this *Server) Handler(conn net.Conn) {
	// 1. new user and insert user into online map
	user := NewUser(conn)
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	// 2. broadcast user online info
	this.Broadcast(user, "is online")

	// 3. block current go
	select {}
}

// Start server
func (this *Server) Start() {
	// 1. socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err :", err)
		return
	}

	// 4. close listen socket
	defer listener.Close()

	// listen publish channel
	go this.ListenPublishChannel()

	for {
		// 2. accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err :", err)
			continue
		}

		// 3. do handler
		go this.Handler(conn)
	}
}