package telegram

import (
	"encoding/json"
	"github.com/klevtcov/makemenu_go/lib"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

//

func newBasePath(token string) string {
	return "bot" + token
}

// функция читающая ответы и обрабатывающая их

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	// метод, добавляющий параметры к запросу
	// Itoa - Integer to ASCII
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	// do request <- getUpdates
	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

// функция по отправке сообщений

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send message: ", err)
	}
	return nil
}

// функция отправляющая запрос и обрабатывающая ответ

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	// defer вызывается всегда в конце родительской функции
	defer func() { err = e.WrapIfErr("can't do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		// path. библиотека сама склеивает части пути добавляя или удаляя слеши
		Path: path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
