package main

import (
	"embed"

	"github.com/yunarsuanto/base-go/http"
	log "github.com/yunarsuanto/base-go/infra/log"
)

//go:embed migrations/*.sql
var EmbedMigration embed.FS

func init() {
	http.EmbedMigration = EmbedMigration
	log.PrintTimestamp()
}

func main() {
	http.ServeHTTP()
}
