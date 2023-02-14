package event_consumer

import (
	"log"
	"time"

	"github.com/klevtcov/makemenu_go/events"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	// бесконечный цикл, который ждёт события и обрабатывает их
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		// если мы получили 0 событий, то ждём 1 секунду
		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		// если событие есть
		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}
	}
}

/*
1. Потеря событий: ретраи, возвращение в хранилище, хранить в памяти и вызывать через какое-то время
отправлять подтвержддени об операции обратно, и только после этого делать сдвиг в процессе
2. обработка всей пачки: если ошибок несколько, то можно останавливаться и ждать какое-то время, потом
пытаться заного обработать всю пачку
3. паралельная обработка: сделать паралельную обработку событий. sync.WaitGroup{}

*/

func (c *Consumer) handleEvents(events []events.Event) error {
	// перебираем события, пытаемся их обработать
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		// если с каким-то из событий произошла ошибка - просто пропускаем его
		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}
	return nil
}
