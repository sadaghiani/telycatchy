package action

import (
	"log"

	"github.com/sadaghiani/telycatchy/observ"
	"github.com/sadaghiani/telycatchy/register"
	"github.com/sadaghiani/telycatchy/storage"
)

type Action struct {
	r       *register.Register
	storage *storage.InMemoryDataStore
	observ  *observ.WatchTower
}

func NewAction(r *register.Register, storage *storage.InMemoryDataStore, observ *observ.WatchTower) *Action {
	return &Action{
		r:       r,
		storage: storage,
		observ:  observ,
	}
}

func (a *Action) Start(key string) *observ.Soldier {

	log.Println("Run Start")

	soldier_1 := observ.NewSolider()
	soldier_1.AddZone(key)

	_, err := a.storage.Get(key)
	if err != nil {
		a.r.Sub(key)
	}

	a.observ.Add(soldier_1)

	return soldier_1
}

func (a *Action) End(soldier_1 *observ.Soldier) {

	a.observ.Remove(soldier_1)

}
