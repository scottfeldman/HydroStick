package main

import (
	"bytes"
	"strconv"
	"time"

	"tinygo.org/x/drivers/net/http"
)

type Blynk struct {
	server string
	token  string
	client http.Client
	traceBuf bytes.Buffer
}

func NewBlynk(token string) *Blynk {
	return &Blynk{
		server: "https://fra1.blynk.cloud",
		token:  token,
		client: http.Client{
			Timeout:     time.Second,
		},
	}
}

func (b *Blynk) updateInt(name string, value int) (err error) {
	url := b.server + "/external/api/update?token=" + b.token + "&" + name + "=" + strconv.Itoa(value)
	res, err := b.client.Get(url)
	if err != nil {
		return err
	} else {
		res.Write(&b.traceBuf)
		trace(b.traceBuf.String())
	}
	return nil
}

func (b *Blynk) event(name string) (err error) {
	url := b.server + "/external/api/logEvent?token=" + b.token + "&code=" + name
	res, err := b.client.Get(url)
	if err != nil {
		return err
	} else {
		res.Write(&b.traceBuf)
		trace(b.traceBuf.String())
	}
	return nil
}
