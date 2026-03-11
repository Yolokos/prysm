package score

var service Service = &MockService{}

func GetService() Service {
	return service
}

func SetService(s Service) {
	service = s
}
