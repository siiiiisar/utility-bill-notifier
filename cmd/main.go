package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gofrs/uuid"
)

func main() {
	token, ok := os.LookupEnv("CHANNEL_ACCESS_TOKEN")
	if !ok {
		panic("Channel Access Token not set")
	}

	retryKey, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	body := map[string]interface{}{
		"messages": []map[string]string{
			{"type": "text", "text": "Github Actionsによって送信されました!"},
		},
	}
	b, err := json.Marshal(body)
	if err != nil {
		panic(err);
	}

	req, err := http.NewRequest("POST", "https://api.line.me/v2/bot/message/broadcast", strings.NewReader(string(b)))
	if err != nil {
		panic(err);
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-Line-Retry-Key", retryKey.String())

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
}
