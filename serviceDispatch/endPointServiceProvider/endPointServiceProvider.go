// endPointServiceProvider project endPointServiceProvider.go
package endPointServiceProvider

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	envHttpApiListenAddr string = "HttpApiListenAddr"
)

var (
	//value exampl  0.0.0.0:8080
	httpApiListenAddr string
)

var (
	errHttpApiListenAddrConf error = errors.New("address not configure or error of http api listen ")
)

func init() {
	initApiListenConf()
}

func initApiListenConf() {
	httpApiListenAddr = os.Getenv(envHttpApiListenAddr)

	log.Println("httpApiListenAddr=", httpApiListenAddr)

	//	_, err := net.ResolveIPAddr("ip", httpApiListenAddr)

	//	if err != nil {
	//		panic(errHttpApiListenAddrConf)
	//	}
}

type EndPointServiceProvider struct {
}

func NewEndPointServiceProvider() *EndPointServiceProvider {
	this := new(EndPointServiceProvider)

	return this
}

func (this *EndPointServiceProvider) ProvideService() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := LoadRouter()

	srv := &http.Server{
		//		Addr: "0.0.0.0:8080",
		Addr: httpApiListenAddr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Minute * 15,
		ReadTimeout:  time.Minute * 15,
		//		ReadHeaderTimeout: time.Minute * 15,
		IdleTimeout: time.Minute * 60,
		//		MaxHeaderBytes:    INT_MAX,
		Handler: router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		} else {
			log.Printf("http server started")
		}
	}()
}
