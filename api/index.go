// package handler

// import (
// 	"net/http"

// 	"github.com/kaitkotak-be/cmd"
// 	"github.com/valyala/fasthttp"
// 	"github.com/valyala/fasthttp/fasthttpadaptor"
// )

// // Handler is the Vercel-compatible serverless function
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	// Initialize Fiber app
// 	app := cmd.NewApp()

// 	// Create a new fasthttp request and response
// 	fctx := &fasthttp.RequestCtx{}

// 	// Convert net/http request to fasthttp request
// 	fasthttpadaptor.ConvertRequest(fctx, r, true)

// 	// Process the request with Fiber
// 	app.Handler()(fctx)

// 	// Copy headers and response body from fasthttp to http.ResponseWriter
// 	fctx.Response.Header.VisitAll(func(k, v []byte) {
// 		w.Header().Set(string(k), string(v))
// 	})
// 	w.WriteHeader(fctx.Response.StatusCode())
// 	w.Write(fctx.Response.Body())
// }
