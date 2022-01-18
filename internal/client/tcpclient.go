/*
	Пакет client создает новое подключение к серверу и содержит интерфейс для работы с ним
*/
package client

import (
	"net"
)

type TcpClient struct {
	conn net.Conn
}

func (t *TcpClient) Write(bytes []byte) (int, error) {
	return t.conn.Write(bytes)
}

func (t *TcpClient) Close() {
	t.conn.Close()
}

func NewTcpClient(port string) (*TcpClient, error) {
	client, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		//log.Printf("Ошибка создания подключения к серверу - %v", err)
		return nil, err
	}
	return &TcpClient{
		conn: client,
	}, nil
}
