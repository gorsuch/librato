package librato

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Gauge struct {
	Name   string  `json:"name"`
	Source string  `json:"source"`
	Count  int64   `json:"count"`
	Sum    float64 `json:"sum"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
}

type Payload struct {
	Gauges []Gauge `json:"gauges"`
}

type Client struct {
	sync.Mutex
	User    string
	Token   string
	payload Payload
}

func (c *Client) AddGauge(g Gauge) {
	c.Lock()
	defer c.Unlock()
	c.payload.Gauges = append(c.payload.Gauges, g)
}

func (c *Client) Flush() error {
	c.Lock()
	defer c.Unlock()

	b, err := json.Marshal(c.payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		"https://metrics-api.librato.com/v1/metrics",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.User, c.Token)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("status code: %d\n", res.StatusCode)
	}

	c.payload = Payload{}

	return nil
}
