package socket

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func tcpSendFile() {

	filePath := ``
	fileSplits := strings.Split(filePath, `\`)
	length := len(fileSplits)
	fileName := fileSplits[length-1]

	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("Dial err: ", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("Conn Write err: ", err)
	}

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Conn read err: ", err)
		return
	}
	if string(buf[:n]) == "ok" {
		sendFile(filePath, conn)
	}
}

func sendFile(filePath string, conn net.Conn) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open File err: ", err)
		return
	}
	defer file.Close()
	buf := make([]byte, 4096)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			fmt.Println("发送文件结束")
			break
		}
		if err != nil {
			fmt.Println("Read File err: ", err)
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Send File err: ", err)
		}
	}
}
