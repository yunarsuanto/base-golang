package main

import (
	"embed"
	"os"

	"github.com/yunarsuanto/base-go/http"
	log "github.com/yunarsuanto/base-go/infra/log"
	"github.com/yunarsuanto/base-go/scripts"
	"github.com/yunarsuanto/base-go/seeds"
)

//go:embed migrations/*.sql
var EmbedMigration embed.FS

func init() {
	http.EmbedMigration = EmbedMigration
	log.PrintTimestamp()
}

func main() {
	http.ServeHTTP()
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		seeds.Seed()
	}

	if len(os.Args) > 1 && os.Args[1] == "crud" {
		name := os.Args[2]
		scripts.GenerateHandlerInput(name)
		scripts.GenerateHandlerInterface(name)
		scripts.GenerateHandler(name)
		scripts.GenerateServiceInterface(name)
		scripts.GenerateService(name)
		scripts.GenerateObject(name)
		scripts.GenerateModel(name)
		scripts.GenerateRepositoryInterface(name)
		scripts.GenerateRepository(name)
		scripts.GenerateRepositoryMap(name)
		return
	}
}
