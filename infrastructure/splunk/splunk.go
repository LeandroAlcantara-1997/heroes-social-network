package splunk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/config"
	outputLog "github.com/LeandroAlcantara-1997/heroes-social-network/ports/output/log"
)

type Splunk struct {
	host   string
	assync bool
}

func New(host string, assync bool) *Splunk {
	return &Splunk{
		host:   host,
		assync: assync,
	}
}

func (s *Splunk) SendErrorLog(ctx context.Context, message string) {
	var outputLog = &outputLog.GlobalLog{
		Data: outputLog.LogMetadata{
			LogLevel: "error",
			Message:  message,
		},
	}
	payload, err := json.Marshal(&outputLog)
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
	var outputEvent = &outputLog.GlobalEvent{
		Data: outputLog.EventMetadata{
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
	url := fmt.Sprintf("%s/services/collector/event", config.Env.SplunkHost)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Default().Printf("error to create request: %v", err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Splunk %s", config.Env.SplunkToken))
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
