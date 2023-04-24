package main

import "net"

type User struct {
	Name 	string
	Addr 	string
	Channel chan string
	conn 	net.Conn
}

// Create a user
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User {
		Name: userAddr,
		Addr: userAddr,
		Channel: make(chan string),
		conn: conn,
	}

	// start go to listen user channel
	go user.ListenMessage()

	return user
}

// Listen user channel and send message to client
func (this *User) ListenMessage() {
	for {
		msg := <- this.Channel
		this.conn.Write([] byte(msg + "\n"))
	}
}

