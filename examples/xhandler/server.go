package main

import (
	"net/http"

	"context"

	"github.com/fzerorubigd/cors"
	"github.com/fzerorubigd/xhandler"
)

func main() {
	c := xhandler.Chain{}

	// Use default options
	c.UseC(cors.Default().HandlerC)

	mux := http.NewServeMux()
	mux.Handle("/", c.Handler(xhandler.HandlerFuncC(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})))

	http.ListenAndServe(":8080", mux)
}
