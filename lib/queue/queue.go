package queue

type Queue interface {
	Consume(func (Msg) error)
	Produce(Msg) error
}

type Msg string
