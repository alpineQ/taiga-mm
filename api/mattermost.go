package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"taiga-mm/config"
	"time"
)

func SendMessage(channel, message string) error {
	url := fmt.Sprintf("%s/hooks/%s", config.Config.MattermostServer, config.Config.MattermostToken)
	data := map[string]string{"text": message, "channel": channel}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	_, err = client.Do(req)
	return err
}
