package custom_client

import (
	"log"
	"net"
	"strconv"
)

func New(port int) net.Conn {
	client, err := net.Dial("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Printf("Ошибка создания подключения к серверу - %v", err.Error())
		return nil
	}
	return client
}
