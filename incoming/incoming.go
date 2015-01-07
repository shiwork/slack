package incoming

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type IncomingConf struct {
	WebHookUrl string
}

// https://api.slack.com/docs/attachments

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool `json:"short"`
}

type Attachment struct {
	Fallback string `json:"fallback"`
	Pretext  string `json:"pretext"`
	Color    string `json:"color"`
	Fields   []*Field `json:"fields"`
}

type Payload struct{
	Attachments []*Attachment `json:"attachments"`
}

func Post(conf IncomingConf, payload Payload) error {
	p, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	_, err = http.PostForm(conf.WebHookUrl, url.Values{
		"payload": []string{string(p)},
	})

	if err != nil {
		return err
	}
	return nil
}
