package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn

	flag int // 客户端当前的工作模式
}

var serverIp string
var serverPort int

func init() {
	// ./client -ip 127.0.0.1 -port 8888
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置默认的服务器ip地址")
	flag.IntVar(&serverPort, "port", 8888, "设置默认服务器端口号")
}

// 创建客户端对象
func NewClient(ip string, port int) *Client {
	client := &Client{
		ServerIp:   ip,
		ServerPort: port,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("net.Dial error :", err)
		return nil
	}
	client.conn = conn

	return client
}

// 显示菜单
func (client *Client) menu() bool {
	var flag int

	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 更新用户名")
	fmt.Println("0. 退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("Please input valid option(0 ~ 3).")
		return false
	}
}

// 处理server的回复消息
func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn)
}

// 公聊模式
func (client *Client) PublicChat() {
	var chatMsg string

	fmt.Println(">>>Please input public message, type exit for exiting")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 消息不为空时发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write error :", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>Please input public message, type exit for exiting")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (client *Client) selectUser() {
	sendMsg := "who\n"
	_, err := client.conn.Write([] byte(sendMsg))
	if err != nil {
		fmt.Println("conn write error :", err)
		return
	}
}

// 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMessage string

	client.selectUser()
	fmt.Println(">>>Please select online user, exit for exiting")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>Please input message content, exit for exiting")
		fmt.Scanln(&chatMessage)

		for chatMessage != "exit" {
			// 消息不为空时发送
			if len(chatMessage) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMessage + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn write error :", err)
					break
				}
			}

			chatMessage = ""
			fmt.Println(">>>Please input message content, exit for exiting")
			fmt.Scanln(&chatMessage)
		}

		client.selectUser()
		fmt.Println(">>>Please select online user, exit for exiting")
		fmt.Scanln(&remoteName)
	}
}

// 更新用户名
func (client *Client) UpdateName() bool {
	fmt.Println(">>> Please input user name:")
	fmt.Scanln(&client.Name)

	message := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(message))
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return false
	}

	return true
}

// 客户端主业务
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
			fmt.Println("Please input valid option(0 ~ 3).")
		}

		switch client.flag {
		case 1:
			client.PublicChat()
			break
		case 2:
			client.PrivateChat()
			break
		case 3:
			client.UpdateName()
			break
		case 0:
			fmt.Println("QUIT client...")
			break
		}
	}
}

func main() {
	// 命令行解析
	flag.Parse()

	// 创建客户端对象
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("Client connect to server failed...")
		return
	}
	fmt.Println("Client connect to server success.")

	// 启动Go程用于处理server的回复消息
	go client.DealResponse()

	// 启动客户端业务
	client.Run()
}
