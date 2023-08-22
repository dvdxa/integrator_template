package agenthandler

import (
	"encoding/xml"
	"integrator_template/domain/agent"
	"integrator_template/utils"
	"io"
	"net/http"
)

const (
	TemplatePayment = "TEMPLATE_PAYMENT"
	PreCheck        = "PRE_CHECK"
)

func (h *Handler) ProcessorHandler(w http.ResponseWriter, r *http.Request) {
	var (
		request  agent.Request
		response *agent.Response
		body     []byte
		err      error
	)
	ok := utils.CheckToken(r.Header.Get("Authorization"), h.cfg.Agent.Token)
	if !ok {
		h.log.Errorf("authentication error")
		utils.WriteXMLResponse(w, http.StatusNetworkAuthenticationRequired, agent.Response{Message: "authentication error",
			Status: agent.ERROR})
		return
	}
	defer r.Body.Close()
	body, err = io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("failed to read req body: %v", err)
		utils.WriteXMLResponse(w, http.StatusInternalServerError, agent.Response{Message: err.Error(), Status: agent.INTERNAL_ERROR})
		return
	}
	h.log.Println("request from processor: ", string(body))
	err = xml.Unmarshal(body, &request)
	if err != nil {
		h.log.Errorf("failed to unmarshal request from processor: %v", err)
		utils.WriteXMLResponse(w, http.StatusInternalServerError, agent.Response{Message: err.Error(), Status: agent.INTERNAL_ERROR})
		return
	}

	switch request.Command {

	case PreCheck:
		adapterRequest := utils.ToAdapterRequest(&request)
		response, err := h.service.PreCheck(adapterRequest)
		if err != nil {
			h.log.Errorf("internal error: %v", err)
			utils.WriteXMLResponse(w, http.StatusInternalServerError, agent.Response{Message: err.Error(), Status: agent.INTERNAL_ERROR})
			return
		}
		h.log.Println("response to processor: ", response)
		utils.WriteXMLResponse(w, http.StatusPartialContent, response)

	case TemplatePayment:
		adaptereRequest := utils.ToAdapterRequest(&request)
		response, err = h.service.MakePayment(adaptereRequest)
		if err != nil {
			h.log.Errorf("internal error: %v", err)
			utils.WriteXMLResponse(w, http.StatusInternalServerError, agent.Response{Message: err.Error(), Status: agent.INTERNAL_ERROR})
			return
		}
		h.log.Println("response to processor: ", response)
		utils.WriteXMLResponse(w, http.StatusPartialContent, response)
	}
}
