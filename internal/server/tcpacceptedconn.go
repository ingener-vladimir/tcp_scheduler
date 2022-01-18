/*
	Пакет server создает новый TCP-сервер и обрабатывает входящие соединения
*/
package server

import (
	"net"
)

type TcpAcceptedConn struct {
	conn net.Conn
}

func (t *TcpAcceptedConn) Close() {
	t.conn.Close()
}

func (t *TcpAcceptedConn) Read(minSizeBuf int) ([]byte, error) {
	bytes := make([]byte, minSizeBuf)
	countReadBytes, err := t.conn.Read(bytes)
	return bytes[:countReadBytes], err
}

func (t *TcpServer) AcceptConn() (*TcpAcceptedConn, error) {
	conn, err := t.server.Accept()
	if err != nil {
		return nil, err
	}
	return &TcpAcceptedConn{
		conn: conn,
	}, nil
}
