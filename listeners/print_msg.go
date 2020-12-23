package listeners

import (
	"auth/events"
	"auth/lib/event_listener"
	"errors"
	"fmt"
	"log"
	"reflect"
)

type PrintMsgListener struct {
}

func (l PrintMsgListener) Handle(e event_listener.Event) error {
	i, ok := e.(events.TestEvent)
	if !ok {
		log.Printf("type is %s \n", reflect.TypeOf(e))
		return errors.New("Undefined event type")
	}

	fmt.Println(i.GetMsg())

	return nil
}
