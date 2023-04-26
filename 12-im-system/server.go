package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 用于存储所有在线用户信息
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 用于广播消息
	PubChan chan string
}

// Create a server instance
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map [string] *User),
		PubChan:   make(chan string),
	}

	return server
}

// Server端广播来自客户端的数据至所有的客户端
func (this *Server) ListenPublishChannel() {
	for {
		msg := <-this.PubChan

		this.mapLock.Lock()
		for _, user := range this.OnlineMap {
			user.Channel <- msg
		}
		this.mapLock.Unlock()
	}
}

// Server端将待广播消息写入channel中
func (this *Server) Broadcast(user *User, msg string) {
	pubMsg := "[" + user.Addr + "]" + user.Name + " " + msg
	this.PubChan <- pubMsg
}

// Server端处理客户端的请求
func (this *Server) Handler(conn net.Conn) {
	// 新建User对象并使User上线
	user := NewUser(conn)
	user.server = this
	user.UserOnline()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 处理来自客户端的数据
	go func() {
		buffer := make([]byte, 4096)

		for {
			// 读取来自客户端的数据
			n, err := conn.Read(buffer)
			if err != nil && err != io.EOF {
				fmt.Println("Conn read err:", err)
				return
			}
			if n == 0 {
				user.UserOffline()
				return
			}

			// 调用用户接口处理客户端数据
			msg := string(buffer[:n - 1])
			user.UserHandleMessage(msg)

			// 有用户消息过来表明当前用户是活跃的
			isLive <- true
		}
	}()

	// 阻塞当前Go程
	for {
		select {
		case <- isLive:
			// 当前用户是活跃的，应该重置定时器
		case <- time.After(time.Second * 60 * 10):
			// 已经超时，将当前用户强制下线
			user.SendMessageToCurrentUser("You were taken offline due to timeout.")
			user.UserOffline()
			close(user.Channel)
			conn.Close()
			return
		}
	}
}

// 启动Server
func (this *Server) Start() {
	// 监听服务器的ip和端口
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err :", err)
		return
	}

	// 退出时关闭监听器
	defer listener.Close()

	// 开启Go程用于广播消息
	go this.ListenPublishChannel()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err :", err)
			continue
		}

		// 处理来自客户端的请求
		go this.Handler(conn)
	}
}
