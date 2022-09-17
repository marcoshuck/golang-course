package main

type DataSource interface {
	Get()
}

func NewStrongDM(client StrongDMClient) DataSource {
	return &strongDM{
		client: client,
	}
}

type DataExporter interface {
	Export()
}

func NewFreshServiceExporter() DataExporter {
	return &freshService{}
}

type DataExchanger interface {
	Run()
}

type DataDestination interface {
	DataSource
	DataExporter
}

type service struct {
	source      DataSource
	destination DataDestination
}

func (s *service) Run() {
	dataSources, err := s.source.Get()
	if err != nil {

	}
	freshServiceData := dataSources.ToFreshService()

	originalFreshServiceData, err := s.destination.Get()
	if err != nil {

	}

	data := originalFreshServiceData.Merge(freshServiceData)

	err = s.destination.Export(data)
	if err != nil {

	}

}

func NewDataExchanger(from DataSource, to DataDestination) DataExchanger {
	return &service{
		source:      from,
		destination: to,
	}
}

func main() {

}
