package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func init() {
	proxy := gin.New()
	proxy.Use(gin.Logger(), gin.Recovery(), Proxy())
	http.Handle("/", proxy)
}
