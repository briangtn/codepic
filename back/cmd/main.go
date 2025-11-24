package main

import (
	"os"

	"github.com/briangtn/codepic/internal/api/routes"
	"github.com/briangtn/codepic/internal/platform/database"
	"github.com/gin-gonic/gin"
)

func main() {
	host := os.Getenv("CODEPIC_HOST")
	if host == "" {
		host = ":8080"
	}
	dbFile := os.Getenv("CODEPIC_DB_FILE")
	if dbFile == "" {
		dbFile = "codepic.db"
	}

	db := database.InitDB(dbFile)

	r := gin.New()

	r.SetTrustedProxies(nil)

	routes.RouterSetup(r, db)

	if err := r.Run(host); err != nil {
		panic(err)
	}
}
