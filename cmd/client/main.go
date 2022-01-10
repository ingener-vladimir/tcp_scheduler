package main

import (
	"encoding/json"
	"fmt"
	"github.com/ingener-vladimir/go_practices/http_scheduler/internal/custom_client"
	"github.com/ingener-vladimir/go_practices/http_scheduler/internal/model"
	"log"
	"strconv"
	"sync"
	"time"
)

const PORT = 8090

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go sendingProcess(i+1, wg)
	}
	wg.Wait()
}

func sendingProcess(seq int, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.Tick(time.Second * time.Duration(seq)):
		fmt.Println("Создание клиента", seq)
		if client := custom_client.New(PORT); client != nil {
			defer client.Close()
			var countMessage int
			for {
				p := model.New(seq, 25, "Ivan_"+strconv.Itoa(countMessage), true)
				data, errMarsh := json.Marshal(p)
				if errMarsh != nil {
					log.Fatal("Ошибка маршалинга сообщения")
				}
				fmt.Println("Будет отправлено сообщение:", string(data))
				client.Write(data)
				countMessage++
				time.Sleep(time.Second * (time.Duration(seq)))
			}
		}
	}
}
