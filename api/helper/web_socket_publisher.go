package helper

import (
	"github.com/gorilla/websocket"
	"net/url"
	"taxi-tracker-api/api/config"
)

type WebSocketPublisher struct {
}

func (WebSocketPublisher) Publish(clientId, msg *string) error {
	u := getConnConfig(clientId)
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte(*msg))
	if err != nil {
		return err
	}
	return nil
}

func getConnConfig(clientId *string) url.URL {
	switch config.ApiEnv {
	case config.Production:
		return url.URL{
			Scheme: "ws",
			Host:   "localhost:3001",
			Path:   "/ws/" + *clientId,
		}
	case config.Local:
		return url.URL{
			Scheme: "ws",
			Host:   "localhost:3001",
			Path:   "/ws/" + *clientId,
		}
	}
	return url.URL{}
}
