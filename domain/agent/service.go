package agent

type AgentService interface {
	MakePayment(adapterRequest *AdapterRequest) (*Response, error)
	PreCheck(adapterRequest *AdapterRequest) (*Response, error)
}

type Service struct {
	adapter AdapterInterface
}

func NewService(adapter AdapterInterface) *Service {
	return &Service{
		adapter: adapter,
	}
}

func (s *Service) PreCheck(adapterRequest *AdapterRequest) (*Response, error) {
	return s.adapter.PreCheck(adapterRequest)
}

func (s *Service) MakePayment(adapterRequest *AdapterRequest) (*Response, error) {
	return s.adapter.TemplatePayment(adapterRequest)
}
