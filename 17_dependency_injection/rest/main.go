package main

func main() {
	svc := NewServiceNotCool()
	server := NewServer(svc)
}

func NewServiceNotCool() Service {
	return &serviceNotSoAwesome{}
}

func NewService() Service {
	return &serviceAwesome{}
}

type serviceNotSoAwesome struct {
}

func (s *serviceNotSoAwesome) Create() {
	//TODO implement me
	panic("implement me")
}

func (s *serviceNotSoAwesome) Read() {
	//TODO implement me
	panic("implement me")
}

func (s *serviceNotSoAwesome) Update() {
	//TODO implement me
	panic("implement me")
}

func (s *serviceNotSoAwesome) Delete() {
	//TODO implement me
	panic("implement me")
}

type serviceAwesome struct {
}

func (s *serviceAwesome) Create() {
	//TODO implement me
	panic("implement me")
}

func (s *serviceAwesome) Read() {
	//TODO implement me
	panic("implement me")
}

func (s *serviceAwesome) Update() {
	//TODO implement me
	panic("implement me")
}

func (s *serviceAwesome) Delete() {
	//TODO implement me
	panic("implement me")
}

type Service interface {
	Create()
	Read()
	Update()
	Delete()
}

type Server struct {
	Service Service
}

func NewServer(svc Service) Server {
	return Server{
		Service: svc,
	}
}
