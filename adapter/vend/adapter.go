package vend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"integrator_template/config"
	"integrator_template/domain/agent"
	"integrator_template/pkg/logger"
	"integrator_template/utils"
	"net/http"
	"time"
)

type Adapter struct {
	cfg        *config.Adapter
	logger     *logger.Logger
	httpClient *http.Client
}

func NewAdapter(cfg *config.Adapter, log *logger.Logger) *Adapter {
	httpClient := &http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}
	return &Adapter{cfg: cfg, logger: log, httpClient: httpClient}
}

func (a Adapter) PreCheck(adapterRequest *agent.AdapterRequest) (*agent.Response, error) {
	var response *agent.AdapterResponse
	//if sign required implement your own logic
	sign, err := utils.CreateSign(adapterRequest, a.cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create sign %v", err)
	}
	jsonData, err := json.Marshal(adapterRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshar adapter request: %v", err)
	}
	resp, err := a.makeRequest(jsonData, agent.PRE_CHECK, sign)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode adapter response: %v", err)
	}
	//if your logic depends on response status, implement own logic
	//if response.Status == 400 {
	//...
	//}
	agentResp := utils.ToAgentResponse(response)
	return agentResp, nil
}

func (a Adapter) TemplatePayment(adapterRequest *agent.AdapterRequest) (*agent.Response, error) {
	var response *agent.AdapterResponse
	sign, err := utils.CreateSign(adapterRequest, a.cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create sign %v", err)
	}
	jsonData, err := json.Marshal(adapterRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshar adapter request: %v", err)
	}
	resp, err := a.makeRequest(jsonData, agent.TEMPLATE_PAYMENT, sign)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode adapter response: %v", err)
	}
	//if your logic depends on response status, implement own logic
	//if response.Status == 400 {
	//...
	//}
	agentResp := utils.ToAgentResponse(response)
	return agentResp, nil
}

func (a Adapter) makeRequest(reqBody []byte, reqAddress string, sign string) (*http.Response, error) {
	headers := make(http.Header)
	headers.Set("Login", a.cfg.Login)
	headers.Set("Password", a.cfg.Password)
	headers.Set("Sign", sign)

	req, err := http.NewRequest(http.MethodPost, a.cfg.URL+reqAddress, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request to adapter: %v", err)
	}

	req.Header = headers

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to adapter: %v", err)
	}
	defer resp.Body.Close()
	return resp, nil
}
