package seeds

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
)

func Seed() {

	db, err := sql.Open("postgres", "user=postgres password=balakutak dbname=umkm sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	seedDir := "./seeds"

	files, err := os.ReadDir(seedDir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".sql") {
			path := filepath.Join(seedDir, f.Name())
			fmt.Printf("Running seed: %s\n", f.Name())
			content, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			if _, err := db.Exec(string(content)); err != nil {
				fmt.Printf("❌ Error in %s: %v\n", f.Name(), err)
			} else {
				fmt.Printf("✅ Seed %s executed successfully\n", f.Name())
			}
		}
	}
}
