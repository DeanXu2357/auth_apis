package listeners

import (
	"auth/lib/email"
	"auth/lib/event_listener"
	"errors"
	"log"
	"reflect"
)

type SendMailListener struct {
}

type SendMailEvent interface {
	To() string
	GetSubject() string
	GetBody() string
}

func (l SendMailListener) Handle(e event_listener.Event) error {
	i, ok := e.(SendMailEvent)
	if !ok {
		log.Printf("type is %s \n", reflect.TypeOf(e))
		return errors.New("Undefined event type")
	}

	info := email.NewInfo()
	err := email.NewEmail(info).SendMail(
		[]string{i.To()},
		i.GetSubject(),
		i.GetBody())
	if err != nil {
		return err
	}

	return nil
}
