// serviceDispatchHost project main.go
package main

import (
	"library/goservice"

	"serviceDispatch/serviceDispatchService"
)

func main() {

	dispatchService := serviceDispatchService.NewServiceDispatchService()

	service := goservice.NewGoService(nil, dispatchService.Start, dispatchService.Stop)

	service.Run()
}
