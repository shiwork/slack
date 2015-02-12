package incoming

import (
    "encoding/json"
    "net/http"
    "net/url"
)

// Config struct
type Config struct {
    WebHookURL string
}

// https://api.slack.com/docs/attachments

// Field struct
type Field struct {
    Title string `json:"title,omitempty"`
    Value string `json:"value,omitempty"`
    Short bool   `json:"short,omitempty"`
}

// Attachment struct
type Attachment struct {
    Fallback string  `json:"fallback,omitempty"`
    Pretext  string  `json:"pretext,omitempty"`
    Color    string  `json:"color,omitempty"`
    Fields   []Field `json:"fields,omitempty"`
}

// Payload struct
type Payload struct {
    Attachments []Attachment `json:"attachments"`
}

// Post method
func Post(conf Config, payload Payload) error {
    p, err := json.Marshal(&payload)
    if err != nil {
        return err
    }

    _, err = http.PostForm(conf.WebHookURL, url.Values{
        "payload": []string{string(p)},
    })

    if err != nil {
        return err
    }
    return nil
}
