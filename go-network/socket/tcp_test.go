package socket

import (
	"testing"
)

func TestCSServer(t *testing.T) {
	tcpCSServer()
}

func TestCSClient(t *testing.T) {
	tcpCSClient()
}

func TestTcpRecvFile(t *testing.T) {
	tcpRecvFile()
}
func TestTcpSendFile(t *testing.T) {
	tcpSendFile()
}
