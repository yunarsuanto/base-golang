package scripts

import (
	"fmt"
	"os"
	"path/filepath"
)

func save(path, code string) {

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		fmt.Println("❌ Error create dir:", err)
	}
	if err := os.WriteFile(path, []byte(code), 0644); err != nil {
		fmt.Println("❌ Error writing:", err)
		return
	}
	fmt.Println("✅ Generated:", path)
}
