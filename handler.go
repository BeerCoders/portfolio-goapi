package main

import (
	"net/http"
	"time"

	env "github.com/BeerCOders/porfolio-go-api/enviroment"
	"github.com/vardius/goserver"
	"golang.org/x/net/context"
)

type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request, *goserver.Context)

func NewHandler(h HandlerFunc) goserver.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c *goserver.Context) {
		start := time.Now()
		ctx, cancel, err := newContext(r)
		if err != nil {
			panic(err)
		}
		defer cancel()
		h(ctx, w, r, c)
		env.Log.Info(ctx, "%s\t%s\t%d", r.Method, r.RequestURI, time.Since(start))
	}
}
