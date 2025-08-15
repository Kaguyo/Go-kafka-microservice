package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/your/project/internal/config"
)

func sendToTarget(cfg config.Config, payload map[string]any) {
	b, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPatch, cfg.TargetServiceURL, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		sendDLQ(cfg, b)
	}
}
