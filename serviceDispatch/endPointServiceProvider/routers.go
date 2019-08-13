package endPointServiceProvider

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const (
	urlRootDir string = "/"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = Routes{
	//example
	Route{
		"Index",
		"GET",
		urlRootDir,
		Index,
	},
	//service detail
	Route{
		"ServiceDetail",
		"GET",
		strings.Join([]string{urlRootDir, "ServiceDetail"}, ""),
		//		"/SDS/ServiceDetail",
		ServiceDetail,
	},
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

type Routes []Route

func LoadRouter() *mux.Router {
	router := mux.NewRouter() //.StrictSlash(true)

	//使用静态文件访问代码行，导致routes配置路由不可用（2018-08-31）
	//	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./api/")))

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
