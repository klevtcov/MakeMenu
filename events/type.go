package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	// текст и тип могут быть в любом мессенджере, а другие поля могут быть разные
	// для тг реализация будет следующая (telegram.go.Meta)
	Meta interface{}
}
