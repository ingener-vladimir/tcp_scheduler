package storage

import (
	"github.com/ingener-vladimir/tcp_scheduler/internal/model"
	"sync"
	"time"
)

type PersonsStorage struct {
	sync.Mutex
	persons map[model.Person]time.Time
}

func NewPersonsStorage() *PersonsStorage {
	return &PersonsStorage{
		persons: make(map[model.Person]time.Time),
	}
}

func (p *PersonsStorage) Add(person model.Person) {
	p.Lock()
	p.persons[person] = time.Now()
	p.Unlock()
}

func (p *PersonsStorage) TakeTime(person model.Person) time.Time {
	p.Lock()
	defer p.Unlock()
	return p.persons[person]
}
