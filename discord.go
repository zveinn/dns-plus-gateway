package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func SendMessageToDiscord(msg string) {
	newE := new(DH)
	newE.Content = msg
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(newE)
	req, _ := http.NewRequest("POST", os.Getenv("ERROR-URL"), payloadBuf)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	_, e := client.Do(req)
	if e != nil {
		log.Println(e)
	}
}

func SendSupportMessageToDiscord(msg string) {
	newE := new(DH)
	newE.Content = msg + " <@913781941361315870>"
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(newE)
	req, _ := http.NewRequest("POST", os.Getenv("ERROR-URL"), payloadBuf)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	_, e := client.Do(req)
	if e != nil {
		log.Println(e)
	}
}

var DISCORDTYPEERROR = "error"

type DH struct {
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Content   string   `json:"content"`
	Embeds    []*Embed `json:"embeds"`
}
type Embed struct {
	Author struct {
		Name    string `json:"name"`
		URL     string `json:"url"`
		IconURL string `json:"icon_url"`
	} `json:"author"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Color       int    `json:"color"`
	Fields      []struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Inline bool   `json:"inline,omitempty"`
	} `json:"fields"`
	Thumbnail struct {
		URL string `json:"url"`
	} `json:"thumbnail"`
	Image struct {
		URL string `json:"url"`
	} `json:"image"`
	Footer struct {
		Text    string `json:"text"`
		IconURL string `json:"icon_url"`
	} `json:"footer"`
}
