package telegram

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strings"

	e "github.com/klevtcov/makemenu_go/lib"
	"github.com/klevtcov/makemenu_go/storage"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	// убираем пробелы из полученного текста
	text = strings.TrimSpace(text)

	log.Printf("got new command `%s` from `%s`", text, username)

	// сохраниение страницы add page: http://...
	// rnd page: /rnd
	// help: /help
	// start: /start: hi + help

	// if isAddCmd(text) {
	// 	// TODO: addPage()
	// 	return p.savePage(chatID, text, username)
	// }

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}

}

func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: can't send random", err) }()

	// пытаемся получить случайное блюдо
	dish, err := p.storage.PickRandomDish(context.Background(), 1)
	if err != nil && !errors.Is(err, storage.ErrNoIngridients) {
		return err
	}

	// если ошибка есть, отправляем сообщение пользователю
	if errors.Is(err, storage.ErrNoIngridients) {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	// пытаемся отправить пользователю блюдо
	if err := p.tg.SendMessage(chatID, dish.TakeDish()); err != nil {
		return err
	}

	return nil
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func isAddCmd(text string) bool {
	return isUrl(text)
}

func isUrl(text string) bool {
	// не распарсит ссылки без протокола, т.е. вида ya.ru
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}

// func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
// 	// в возврате ошибка именованная, т.к. ожидается в разных местах
// 	// сразу задаём отложенную обработку ошибок через defer
// 	defer func() { err = e.WrapIfErr("can't do command: save page", err) }()

// 	dish := &storage.Page{
// 		URL:      pageURL,
// 		UserName: username,
// 	}

// 	// проверяем, существует ли у нас данная страница в сохранённых
// 	isExists, err := p.storage.IsExist(context.Background(), page)
// 	if err != nil {
// 		return err
// 	}

// 	//  если существует, возвращаем в чат сообщение
// 	if isExists {
// 		return p.tg.SendMessage(chatID, msgAlreadyExists)
// 	}

// 	// пытаемся сохранить страницу
// 	if err := p.storage.Save(context.Background(), page); err != nil {
// 		return err
// 	}

// 	// если сохранилась корректно, отправляем сообщение пользователю
// 	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
// 		return err
// 	}

// 	return nil
// }
