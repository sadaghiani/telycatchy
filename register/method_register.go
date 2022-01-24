package register

import (
	"fmt"
	"log"
	"reflect"

	"github.com/sadaghiani/telycatchy/observ"
	"github.com/sadaghiani/telycatchy/storage"
)

type Register struct {
	rcvr    interface{}
	storage *storage.InMemoryDataStore
	observ  *observ.WatchTower
}

func NewRegister(s *storage.InMemoryDataStore, observ *observ.WatchTower) *Register {
	return &Register{
		storage: s,
		observ:  observ,
	}
}

func (r *Register) Set(rcvr interface{}) {
	r.rcvr = rcvr
}

func (r *Register) Connect() {
	inputValue := reflect.ValueOf(r.rcvr).MethodByName("Connect").Call(nil)
	// rc := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, reflect.TypeOf(inputValue)), 0)
	if inputValue[0].Kind() != reflect.Chan {
		panic(fmt.Sprintf("DropChan called on invalid type: %s", inputValue[0].Type()))
	}

	go func() {
		for {
			item, ok := inputValue[0].Recv()
			if !ok {
				break
			}

			r.storage.Set(item.MapKeys()[0].String(), storage.NewResultContent(item.MapIndex(item.MapKeys()[0]).Interface()))
			r.observ.Notify(item.MapKeys()[0].String())
		}
	}()
}

func (r *Register) Sub(param1 string) {
	// m := map[string][]string{"foo": []string{"bar"}}

	in := []reflect.Value{reflect.ValueOf(param1)}
	reflect.ValueOf(r.rcvr).MethodByName("Sub").Call(in)
	log.Println("Sub.", param1)

}

func (r *Register) UnSub(param1 string) {
	// m := map[string][]string{"foo": []string{"bar"}}

	in := []reflect.Value{reflect.ValueOf(param1)}
	reflect.ValueOf(r.rcvr).MethodByName("UnSub").Call(in)
	log.Println("Unsub.", param1)

}
