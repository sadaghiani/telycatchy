package observ

import (
	"sync"

	"github.com/sadaghiani/telycatchy/storage"

	"github.com/google/uuid"
)

type (
	Observable interface {
		Add(observer Observer)
		Notify(event interface{})
		Remove(event interface{})
	}

	Observer interface {
		NotifyCallback(event interface{}, wt *WatchTower)
	}

	WatchTower struct {
		Observer sync.Map
		// Data     map[string]string
		Data *storage.InMemoryDataStore
	}

	Soldier struct {
		Id   int
		Zone string
		C    chan interface{}
	}
)

func NewObserv(s *storage.InMemoryDataStore) *WatchTower {
	return &WatchTower{
		Data: s,
	}
}

func (wt *WatchTower) Add(observer Observer) {
	wt.Observer.Store(observer, struct{}{})
}

func (wt *WatchTower) Remove(observer Observer) {
	wt.Observer.Delete(observer)
}

func (wt *WatchTower) Notify(event interface{}) {
	wt.Observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}

		key.(Observer).NotifyCallback(event, wt)
		return true
	})
}

func (s *Soldier) NotifyCallback(event interface{}, wt *WatchTower) {
	if event.(string) == s.Zone {
		res, _ := wt.Data.Get(event.(string))
		s.C <- res.Content.Data
	}
}

func NewSolider() *Soldier {
	return &Soldier{
		Id: int(uuid.New().ID()),
		C:  make(chan interface{}),
	}
}

func (s *Soldier) AddZone(zone string) {
	s.Zone = zone
}
