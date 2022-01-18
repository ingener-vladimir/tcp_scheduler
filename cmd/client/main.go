/*
	Через равный промежуток времени (1 сек) создается TCP-клиент, который подключается к серверу по порту из аргументов,
	и выполняется отправка на сервер сообщений
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ingener-vladimir/tcp_scheduler/internal/client"
	"github.com/ingener-vladimir/tcp_scheduler/internal/model"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// gracefully shutdown (system call)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sendingProcess(port, i+1, wg)
	}
	wg.Wait()
}

func sendingProcess(port string, seq int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-time.Tick(time.Second * time.Duration(seq))
	log.Println("Создание клиента", seq)
	client, err := client.NewTcpClient(port)
	if err != nil {
		log.Printf("Ошибка создания нового подключения %d: %v\n", seq, err)
		return
	}

	defer client.Close()
	var countMessage int
	for {
		p := model.New(seq, 25, "Ivan_"+strconv.Itoa(countMessage), true)
		data, errMarsh := json.Marshal(p)
		if errMarsh != nil {
			log.Printf("Ошибка маршалинга сообщения %v: %v", p, err)
			continue
		}
		fmt.Println("Будет отправлено сообщение:", p)
		if _, errWrite := client.Write(data); errWrite != nil {
			log.Println("Ошибка отправки сообщения", errWrite)
			return
		}
		countMessage++
		time.Sleep(time.Second * (time.Duration(seq)))
	}
}
