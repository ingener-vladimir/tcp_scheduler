package server

import (
	"net"
)

type TcpServer struct {
	server net.Listener
}

func NewTcpServer(port string) (*TcpServer, error) {
	server, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		return nil, err
	}
	return &TcpServer{
		server: server,
	}, nil
}
