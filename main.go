package main

import (
	"flag"
	"github.com/valyala/fasthttp"
	"log"
)

var (
	notFoundResponse   = []byte(`<!DOCTYPE html><html><head><meta name=viewport content="width=device-width,initial-scale=1"><title>Sorry, Not found.</title></head><body><h1>404</h1><p>Not Found</p></body></html>`)
	contentTypeHeaderK = []byte("Content-Type")
	contentTypeHeaderV = []byte("text/html;charset=utf8")
	locationHeaderK    = []byte("Location")
	locationHeaderV    []byte
	notFoundAddr       = flag.String("notFoundAddr", ":8088", "TCP address to listen to")
	redirAddr          = flag.String("redirAddr", ":8089", "TCP address to listen to")
	redirDest          = flag.String("redirDest", "http://ruskline.ru", "TCP address to listen to")
)

func init() {
	locationHeaderV = []byte(*redirDest)
}

func llbNotFoundHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(404)
	ctx.Response.Header.SetBytesKV(contentTypeHeaderK, contentTypeHeaderV)
	ctx.Write(notFoundResponse)
}
func llbProdRewriteHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetBytesKV(locationHeaderK, locationHeaderV)
}

func main() {
	flag.Parse()
	go func() {
		if err := fasthttp.ListenAndServe(*notFoundAddr, llbNotFoundHandler); err != nil {
			log.Fatalf("error in ListenAndServe (notFound): %s", err)
		}
	}()
	if err := fasthttp.ListenAndServe(*redirAddr, llbProdRewriteHandler); err != nil {
		log.Fatalf("error in ListenAndServe (redir): %s", err)
	}
}
