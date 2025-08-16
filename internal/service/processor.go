package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-kafka-microservice/internal/config"
	"net/http"
	"time"
)

func SendToTarget(cfg config.Config, payload map[string]any) {

	b, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPatch, cfg.TargetServiceURL, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		SendDLQ(cfg, b)
		fmt.Printf("Erro ao enviar para o serviço alvo: %v, status: %d\n", err, resp.StatusCode)
	}
}
