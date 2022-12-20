package nostrate

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

type Client struct {
	sc *websocket.Conn
}

func NewClient(httpClient *http.Client, relayAddress string) (*Client, error) {
	dialCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	opt := websocket.DialOptions{
		HTTPClient: httpClient,
	}
	c, _, err := websocket.Dial(dialCtx, relayAddress, &opt)
	if err != nil {
		return nil, fmt.Errorf("error connection to relay %w", err)
	}
	return &Client{c}, nil
}

func (c *Client) Close() {
	c.sc.Close(websocket.StatusGoingAway, "client was shut down")
}
