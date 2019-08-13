// endPointOverseeHost project main.go
package main

import (
	"endPointOversee/endPointOverseeService"
	"library/goservice"
)

func main() {
	overseeService := endPointOverseeService.NewEndPointOverseeService()

	service := goservice.NewGoService(nil, overseeService.Start, overseeService.Stop)

	service.Run()
}
