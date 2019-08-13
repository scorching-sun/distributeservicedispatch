// endPointServiceAlarmer project endPointServiceAlarmer.go
package endPointServiceAlarmer

import (
	"serviceDispatch"
)

type EndPointServiceAlarmer struct {
}

func NewEndPointServiceAlarmer() *EndPointServiceAlarmer {
	this := new(EndPointServiceAlarmer)

	return this
}

func (this *EndPointServiceAlarmer) FollowProcess(event *serviceDispatch.Event) error {
	return nil
}
