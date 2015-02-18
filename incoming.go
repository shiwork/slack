package slack

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// https://api.slack.com/docs/attachments

// Field struct
type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

// Attachment struct
type Attachment struct {
	Fallback string `json:"fallback,omitempty"`

	Pretext string `json:"pretext,omitempty"`

	Color string `json:"color,omitempty"`

	AuthorName string `json:"author_name,omitempty"`
	AuthorLink string `json:"author_link,omitempty"`
	AuthorIcon string `json:"author_icon,omitempty"`

	Title     string `json:"title,omitempty"`
	TitleLink string `json:"title_link,omitempty"`

	Text string `json:"text,omitempty"`

	Fields []Field `json:"fields,omitempty"`

	ImageUrl string `json:"image_url,omitempty"`
}

// Payload struct
type Payload struct {
	Attachments []Attachment `json:"attachments"`
}

type Incoming struct {
	WebHookURL string
}

// Post method
func (s *Incoming) Post(payload Payload) error {
	p, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	_, err = http.PostForm(s.WebHookURL, url.Values{
		"payload": []string{string(p)},
	})

	if err != nil {
		return err
	}
	return nil
}
