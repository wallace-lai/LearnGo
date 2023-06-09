package main

import (
	"net"
	"strings"
)

type User struct {
	Name    string
	Addr    string
	Channel chan string
	conn    net.Conn
	server  *Server // User归属的Server对象
}

// 创建一个用户
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		Channel: make(chan string),
		conn:    conn,
	}

	// 启动Go程监听channel
	go user.ListenMessage()

	return user
}

// 监听channel
func (this *User) ListenMessage() {
	for {
		msg := <-this.Channel
		this.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线业务
func (this *User) UserOnline() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	this.server.Broadcast(this, "is online")
}

// 用户下线业务
func (this *User) UserOffline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.Broadcast(this, "is offline")
}

// 给当前User对应的客户端发送消息
func (this *User) SendMessageToCurrentUser(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息业务
func (this *User) UserHandleMessage(msg string) {
	if msg == "who" {
		// 查看当前的在线用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineUser := "[" + user.Addr + "]" + user.Name + " : ONLINE\n"
			this.SendMessageToCurrentUser(onlineUser)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 更新用户名
		newName := strings.Split(msg, "|")[1]

		this.server.mapLock.Lock()
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMessageToCurrentUser("Your new name " + newName + " is occupied, try another name pls.\n")
		} else {
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.Name = newName
			this.SendMessageToCurrentUser("Change name to " + newName + " success.\n")
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 私发消息，消息格式为：to|张三|消息内容

		// 获取接收方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMessageToCurrentUser("Invalid message format.\n")
			return
		}

		// 根据用户名找到接收方user对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMessageToCurrentUser("Remote user does NOT exist.\n")
			return
		}

		// 获取消息内容，通过接收方user对象将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMessageToCurrentUser("Message content is null.\n")
			return
		}
		remoteUser.SendMessageToCurrentUser(this.Name + "talks |" + content)
	} else {
		this.server.Broadcast(this, msg)
	}
}
