package main

import (
	// "context"
	"fmt"
	"github.com/klevtcov/makemenu_go/storage/sqlite"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	// storagePath = "storage"
	sqliteStorage = "data/sqlite/storage.db"
	batchSize     = 100
)

func main() {

	storage, err := sqlite.New(sqliteStorage)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}
	// if err := storage.Init(context.TODO()); err != nil {
	// 	log.Fatal("can't init storage: ", err)
	// }

	// eventsProccessor := telegram.New(tgClient.New(tgBotHost, mustToken()), storage)

	// log.Print("service started")

	// consumer := event_consumer.New(eventsProccessor, eventsProccessor, batchSize)
	// if err := consumer.Start(); err != nil {
	// 	log.Fatal("service is stopped", err)
	// }

	fmt.Println("app started")
}
