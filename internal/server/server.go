package server

type (
	server struct {
		storage storage
	}

	storage interface {
	}
)

type NewServerParams struct {
	Host    string
	Storage storage
}

func New(params NewServerParams) *server {
	return &server{}
}

func (s *server) Start() <-chan error {
	return nil
}

func (s *server) Stop() error {
	return nil
}
