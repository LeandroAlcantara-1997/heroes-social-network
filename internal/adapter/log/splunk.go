package log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	otelhttp "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
)

type splunk struct {
	host    string
	token   string
	assync  bool
	strdout *zap.Logger
}

func New(host, token string, assync bool, zapLogger *zap.Logger) *splunk {
	return &splunk{
		host:    host,
		token:   token,
		assync:  assync,
		strdout: zapLogger,
	}
}

func (s *splunk) Send(ctx context.Context, payload []byte) {
	if s.assync {
		go s.send(ctx, payload)
	} else {
		s.send(ctx, payload)
	}
}

func (s *splunk) send(ctx context.Context, payload []byte) {
	url := fmt.Sprintf("%s/services/collector/event", s.host)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		s.strdout.Error(fmt.Sprintf("error to create request: %s", err.Error()))
	}

	request.Header.Add("Authorization", fmt.Sprintf("Splunk %s", s.token))
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	response, err := client.Do(request)
	if err != nil {
		s.strdout.Error(fmt.Sprintf("error to send request: %s", err.Error()))
		return
	}

	if response.StatusCode != http.StatusOK {
		responseMap := make(map[string]any)
		if err := json.NewDecoder(response.Body).Decode(&responseMap); err != nil {
			s.strdout.Error(fmt.Sprintf("error to decode response body: %s", err.Error()))
			return
		}
		s.strdout.Error(fmt.Sprintf("error to send request: %v", responseMap))
	}
}
