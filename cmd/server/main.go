package main

import (
	"fmt"
	"github.com/ingener-vladimir/tcp_scheduler/internal/model"
	"github.com/ingener-vladimir/tcp_scheduler/internal/server"
	"github.com/ingener-vladimir/tcp_scheduler/internal/storage"
	"github.com/ingener-vladimir/tcp_scheduler/internal/utils"
	"io"
	"log"
	"os"
)

const minSizeBuf = 100

func main() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	server, err := server.NewTcpServer(port)
	if err != nil {
		log.Fatal("Ошибка при создании TCP сервера", err)
	}
	persons := storage.NewPersonsStorage()
	for {
		conn, err := server.AcceptConn()
		if err != nil {
			log.Println("Ошибка обработки нового подключения к TCP серверу", err)
			continue
		}
		go requestProcess(conn, persons)
	}
}

func requestProcess(conn *server.TcpAcceptedConn, persons *storage.PersonsStorage) {
	defer conn.Close()
	for {
		bytes, err := conn.Read(minSizeBuf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Соединение закрыто!")
				break
			}
		}
		var p model.Person
		errMarsh := utils.JsonUnmarshal(bytes, &p)
		if errMarsh != nil {
			return
		}
		fmt.Printf("Новая личность - %v\n", p)
		persons.Add(p)
	}
}
