package main

import (
	"os"

	"github.com/briangtn/codepic/internal/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.SetTrustedProxies(nil)

	routes.RouterSetup(r)

	host := os.Getenv("CODEPIC_HOST")
	if host == "" {
		host = ":8080"
	}

	if err := r.Run(host); err != nil {
		panic(err)
	}
}
