package socket

import (
	"fmt"
	"net"
	"strings"
)

func handlerConnClient(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	fmt.Printf("连接客户端：%v 成功 \n", addr.String())

	// 传输数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Printf("客户端 %v 已关闭 \n", addr)
			break
		}
		if err != nil {
			fmt.Println("读取传输数据失败：", err)
			return
		}
		if string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit\r\n" {
			fmt.Printf("检测到客户端：%v 退出\n", addr)
			break
		}
		fmt.Printf("服务器读到 %v 客户端发送数据：%s \n", addr, string(buf[:n]))

		// _, err = conn.Write([]byte("hello, i am server"))
		// if err != nil {
		// 	fmt.Println("写入传输数据失败：", err)
		// 	return
		// }
		// 处理数据: 小 -- 大
		UpperWord := strings.ToUpper(string(buf[:n]))
		// 写
		conn.Write([]byte(UpperWord))
	}
}

func tcpCSServer() {
	// 创建一个 socket 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:7020")
	if err != nil {
		fmt.Println("创建 socket 失败：", err)
		return
	}
	defer listener.Close()
	fmt.Println("监听套接字，创建成功。。。")

	for {
		// 等待阻塞 获取网络连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("等待阻塞，创建网络连接失败：", err)
			return
		}
		go handlerConnClient(conn)
	}
}
