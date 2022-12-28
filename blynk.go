package main

import (
	"strconv"
	"time"

	"tinygo.org/x/drivers/net"
)

type Blynk struct {
	server string
	token  string
	client HttpClient
}

func NewBlynk(token string) *Blynk {
	return &Blynk{
		server: "https://fra1.blynk.cloud",
		token:  token,
		client: HttpClient{
			timeout:     time.Second,
			connections: map[string]net.Conn{},
		},
	}
}

func (b *Blynk) updateInt(name string, value int) (err error) {
	url := b.server + "/external/api/update?token=" + b.token + "&" + name + "=" + strconv.Itoa(value)
	req := NewGET(url, nil)
	res, err := b.client.sendHttp(req, false)
	if err != nil {
		return err
	} else {
		println(string(res.bytes))
	}
	return nil
}

func (b *Blynk) event(name string) (err error) {
	url := b.server + "/external/api/logEvent?token=" + b.token + "&code=" + name
	req := NewGET(url, nil)
	res, err := b.client.sendHttp(req, false)
	if err != nil {
		return err
	} else {
		println(string(res.bytes))
	}
	return nil
}
