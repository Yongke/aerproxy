package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
)

var once = &sync.Once{}
var backend string

func getBackend() {
	backend = os.Getenv("PROXY_PASS_HOST_PORT")
}

func Proxy() gin.HandlerFunc {
	once.Do(getBackend)
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = backend
		}
		proxy := &ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
