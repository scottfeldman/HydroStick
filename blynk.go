package main

import "fmt"

type Blynk struct {
	token  string
	client HttpClient
}

func NewBlynk(token string) *Blynk {
	return &Blynk{
		token: token,
		client: HttpClient{
			server: "blynk.cloud",
		},
	}
}

func (b *Blynk) updateInt(name string, value int) {
	path := fmt.Sprintf("/external/api/update?%s=%d&token=%s", name, value, b.token)
	resp, err := b.client.request(path)
	if err != nil {
		println(err.Error())
	} else {
		println(string(resp))
	}
}
