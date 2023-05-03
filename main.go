package main

import (
	// "fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgClient "github.com/klevtcov/makemenu_go/clients/telegram"
	event_consumer "github.com/klevtcov/makemenu_go/consumer/event-consumer"
	"github.com/klevtcov/makemenu_go/events/telegram"
	"github.com/klevtcov/makemenu_go/storage/sqlite"
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

	eventsProccessor := telegram.New(tgClient.New(tgBotHost, tgToken), storage)

	// fmt.Println("app started")

	consumer := event_consumer.New(eventsProccessor, eventsProccessor, batchSize)

	go func() {
		if err := consumer.Start(); err != nil {
			log.Fatal("service is stopped", err)
		}
	}()

	log.Print("service started")

	// Ожидаем сигнала на закрытие приложения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := sqlite.Shutdown(storage); err != nil {
		log.Fatalf("error occured on db connection close: %s", err.Error())
	}

	log.Print("service stopped")
}
