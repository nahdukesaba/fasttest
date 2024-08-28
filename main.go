package main

import (
	"github.com/gin-gonic/gin"

	"fasttest/server"
)

func main() {
	r := gin.Default()
	s := server.NewServer(r)
	s.RegisterRoute()

	s.Run()
}
