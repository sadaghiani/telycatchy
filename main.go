package telycatchy

import (
	"github.com/sadaghiani/telycatchy/action"
	"github.com/sadaghiani/telycatchy/observ"
	"github.com/sadaghiani/telycatchy/register"
	"github.com/sadaghiani/telycatchy/storage"
)

func NewAction(i interface{}) *action.Action {
	s := storage.NewInMemoryDataStore()
	o := observ.NewObserv(s)
	r := register.NewRegister(s, o)

	r.Set(i)
	r.Connect()

	return action.NewAction(r, s, o)
}
