package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateObject(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package objects

		type Create%sRequest struct {
			Name string
		}

		type Update%sRequest struct {
			Id       string
			Name string
		}

		type Delete%sRequest struct {
			Id string
		}

		type List%sResponse struct {
			Id       string
			Name string
		}

	`,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
	)

	code := strings.ReplaceAll(template, "__BACKTICK__", "`")
	filePath := fmt.Sprintf("objects/%s_object.go", name)

	save(filePath, code)
}
