package log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Splunk struct {
	host   string
	token  string
	assync bool
}

func New(host string, token string, assync bool) *Splunk {
	return &Splunk{
		host:   host,
		token:  token,
		assync: assync,
	}
}

func (s *Splunk) SendErrorLog(ctx context.Context, err error) {
	logger := GetLogger(ctx)
	logger.Data.LogLevel = "error"
	logger.Data.Message = err.Error()

	payload, err := json.Marshal(&logger)
	if err != nil {
		log.Fatalf("Error to do marshal log: %v", err)
	}
	if s.assync {
		go s.sendToSplunk(ctx, payload)
	} else {
		s.sendToSplunk(ctx, payload)
	}
}

func (s *Splunk) SendEvent(ctx context.Context, eventCode int, message string) {
	var outputEvent = &GlobalEvent{
		Data: &EventMetadata{
			EventCode: eventCode,
			Message:   message,
		},
	}
	payload, err := json.Marshal(&outputEvent)
	if err != nil {
		log.Fatalf("Error to do marshal log: %v", err)
	}
	if s.assync {
		go s.sendToSplunk(ctx, payload)
	} else {
		s.sendToSplunk(ctx, payload)
	}
}

func (s *Splunk) sendToSplunk(ctx context.Context, payload []byte) {
	url := fmt.Sprintf("%s/services/collector/event", s.host)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Default().Printf("error to create request: %v", err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Splunk %s", s.token))
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		log.Default().Printf("error to send request: %v", err)
		return
	}

	if response.StatusCode != http.StatusOK {
		responseMap := make(map[string]any)
		if err := json.NewDecoder(response.Body).Decode(&responseMap); err != nil {
			log.Default().Printf("error to decode response body: %v", err)
			return
		}
		log.Default().Printf("error to send request: %v", responseMap)
		return
	}
}
