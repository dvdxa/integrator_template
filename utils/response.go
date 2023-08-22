package utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"integrator_template/domain/agent"
	"net/http"
)

func WriteXMLResponse(w http.ResponseWriter, httpStatus int, v interface{}) error {
	body, err := xml.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal to XML: %v", err)
	}
	w.WriteHeader(httpStatus)
	_, err = w.Write(body)
	if err != nil {
		return fmt.Errorf("failed to write data to connection: %v", err)
	}
	return nil
}

func WriteJSONResponse(w http.ResponseWriter, httpStatus int, v interface{}) error {
	body, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON: %v", err)
	}
	w.WriteHeader(httpStatus)
	_, err = w.Write(body)
	if err != nil {
		return fmt.Errorf("failed to write data to connection: %v", err)
	}
	return nil
}

func ToAgentResponse(adapterResponse *agent.AdapterResponse) *agent.Response {
	var response *agent.Response
	response.Status = int64(adapterResponse.Status)
	response.Message = adapterResponse.Message
	//if there are other fields, implement your own logic
	return response
}

func ToAdapterRequest(humoRequest *agent.Request) *agent.AdapterRequest {
	var request *agent.AdapterRequest
	request.Amount = humoRequest.Payment.Amount
	//implement own logic
	return request
}
