package socket

import (
	"fmt"
	"net"
	"os"
)

func tcpRecvFile() {
	listener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("Create socket err: ", err)
		return
	}

	defer listener.Close()
	fmt.Println("监听套接字，创建成功")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Conn Accept err: ", err)
		return
	}

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	fileName := string(buf[:n])
	fmt.Println("准备要接受的文件名：", fileName)

	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("Conn Write err: ", err)
		return
	}

	recvFile(fileName, conn)
}

func recvFile(fileName string, conn net.Conn) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create file err: ", err)
		return
	}

	defer file.Close()
	buf := make([]byte, 0)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("读取传输文件结束")
			break
		}
		if err != nil {
			fmt.Println("读取传输文件出错：", err)
			return
		}
		file.Write(buf[:n])
	}
}
