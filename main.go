package main

import (
	"flag"
	"fmt"
	"net/http"

	//	env "github.com/BeerCOders/porfolio-go-api/enviroment"
	"github.com/EconomistDigitalSolutions/ramlapi"
	"github.com/vardius/goserver"
	"golang.org/x/net/context"
)

var (
	programPort = flag.String("port", "8080", "Set server port, default 8080")
	RouteMap    = map[string]goserver.HandlerFunc{
		"Root":    Root,
		"Version": Version,
	}
)

func main() {
	flag.Parse()

	api, _ := ramlapi.ProcessRAML("api.raml")
	ramlapi.Build(api, routerFunc)

	//	env.Log.Critical(context.TODO(), "%s", http.ListenAndServe(":"+*programPort, env.Server))
}

func routerFunc(ep *ramlapi.Endpoint) {
	switch ep.Verb {
	case "GET":
		router.GET(ep.Path, RouteMap[ep.Handler])
	}
}

func Version(_ context.Context, w http.ResponseWriter, _ *http.Request, _ *goapi.Context) {
	fmt.Fprint(w, "VERSION\n")
}

func Root(_ context.Context, w http.ResponseWriter, _ *http.Request, _ *goapi.Context) {
	fmt.Fprint(w, "HOME\n")
}
