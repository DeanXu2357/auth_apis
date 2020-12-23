package events

type TestEvent struct {
}

func (e TestEvent) GetName() string {
	return Test
}

func (e TestEvent) GetMsg() string{
	return "this is test msg event"
}
