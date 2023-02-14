package telegram

import (
	"errors"

	"github.com/klevtcov/makemenu_go/clients/telegram"
	"github.com/klevtcov/makemenu_go/events"
	e "github.com/klevtcov/makemenu_go/lib"
	"github.com/klevtcov/makemenu_go/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	Username string
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	// получаем апдейты
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events: ", err)
	}

	// если список апдейтов пустой - заканчиваем работу функции
	if len(updates) == 0 {
		return nil, nil
	}

	// готовим переменную для результата и аллоцируем память под неё
	// срез, длинна 0, капасити = длинне массива апдейтов
	res := make([]events.Event, 0, len(updates))

	// перебираем все апдейты и преобразуем их в тип event
	for _, u := range updates {
		res = append(res, event(u))
	}

	// обновляем параметр offset, чтобы в следующий раз получить новые апдейты
	p.offset = updates[len(updates)-1].ID + 1

	// возвращаем результат
	return res, nil
}

// выполняет различные действия, в зависимости от типа ивенда

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("can't process message: ", ErrUnknownEventType)
	}
}

func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message: ", err)

	}

	// в зависимости от типа сообщения, нам надо выполнить разные действия
	if err := p.doCmd(event.Text, meta.ChatID, meta.Username); err != nil {
		return e.Wrap("can't process message: ", err)
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("can't get meta: ", ErrUnknownMetaType)
	}

	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}
	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}
