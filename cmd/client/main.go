package main

import (
	"encoding/json"
	"fmt"
	"github.com/ingener-vladimir/go_practices/http_scheduler/cmd/internal/model"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		fmt.Println("Создание клиента", i+1)
		//client, err := net.Dial("TCP", "localhost:8000")
		//if err != nil {
		//	log.Println("Ошибка создания подключения к серверу клиентом", i+1)
		//	continue
		//}
		wg.Add(1)
		go connect(i+1, wg, nil)
	}
	wg.Wait()
}

func connect(seq int, wg *sync.WaitGroup, conn net.Conn) {
	//defer conn.Close()
	defer wg.Done()

	p := model.New(seq, 25, "Ivan", true)
	data, errMarsh := json.Marshal(p)
	if errMarsh != nil {
		log.Fatal("Ошибка маршалинга сообщения")
	}
	for {
		select {
		case <-time.Tick(time.Second * time.Duration(seq)):
			fmt.Println(seq, data)
			//conn.Write(data)
		}
	}
}
