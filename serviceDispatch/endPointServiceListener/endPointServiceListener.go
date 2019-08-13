// endPointServiceListener project endPointServiceListener.go
package endPointServiceListener

import (
	"errors"
	"net"
	"os"
	"serviceDispatch"
	"strconv"
)

const (
	envKeySDSListenIP   string = "SDSListenIP"
	envKeySDSListenPort string = "SDSListenPort"
)

var (
	SDSListenIP   net.IP
	SDSListenPort int

	errSDSListenIPNil   error = errors.New("sds listen ip invaild")
	errSDSListenPortNil error = errors.New("sds listen port invaild")
)

func init() {
	ip := os.Getenv(envKeySDSListenIP)

	SDSListenIP = net.ParseIP(ip)

	if SDSListenIP == nil {
		panic(errSDSListenIPNil)
	}

	port := os.Getenv(envKeySDSListenPort)

	var err error
	SDSListenPort, err = strconv.Atoi(port)

	if err != nil {
		panic(errSDSListenPortNil)
	}
}

type EndPointServiceListener struct {
	Observers map[serviceDispatch.Observer]struct{}
}

func NewEndPointServiceListener() *EndPointServiceListener {
	this := new(EndPointServiceListener)

	return this
}

func (this *EndPointServiceListener) Regist(ob serviceDispatch.Observer) {
	this.Observers[ob] = struct{}{}
}

func (this *EndPointServiceListener) Deregist(ob serviceDispatch.Observer) {
	delete(this.Observers, ob)
}

func (this *EndPointServiceListener) Notify(event *serviceDispatch.Event) {
	for ob, _ := range this.Observers {
		ob.FollowProcess(event)
	}
}
