// serviceDispatchService project serviceDispatchService.go
package serviceDispatchService

type ServiceDispatchService struct {
}

func NewServiceDispatchService() *ServiceDispatchService {
	this := new(ServiceDispatchService)

	return this
}

//service start
func (this *ServiceDispatchService) Start() error {
	var err error

	defer func() {
		if p := recover(); p != nil {
			if er, ok := p.(error); ok {
				err = er
			}
		}
	}()

	this.dispatchService()

	return err
}

//service stop
func (this *ServiceDispatchService) Stop() error {
	var err error

	defer func() {
		if p := recover(); p != nil {
			if er, ok := p.(error); ok {
				err = er
			}
		}
	}()

	return err
}
