package main

import (
	// "context"
	"fmt"
	tgClient "github.com/klevtcov/makemenu_go/clients/telegram"
	event_consumer "github.com/klevtcov/makemenu_go/consumer/event-consumer"
	"github.com/klevtcov/makemenu_go/events/telegram"
	"github.com/klevtcov/makemenu_go/storage/sqlite"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	// storagePath = "storage"
	sqliteStorage = "data/sqlite/makemenu.db"
	batchSize     = 100
)

func main() {

	storage, err := sqlite.New(sqliteStorage)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	eventsProccessor := telegram.New(tgClient.New(tgBotHost, token), storage)

	fmt.Println("app started")

	// log.Print("service started")

	consumer := event_consumer.New(eventsProccessor, eventsProccessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}
