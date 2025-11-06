package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateRepositoryMap(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package %s_repository

		import (
			"fmt"
			"strings"
		)

		func mapQueryFilterList%s(search string, params *[]any) (string, error) {
			var result string
			var filterArray []string

			if search != "" {
				searchParams := fmt.Sprintf("%%%%s%%", search)
				filterArray = append(filterArray, "u.username LIKE $1")
				(*params) = append((*params), searchParams)
			}

			result = strings.Join(filterArray, " AND ")
			if result != "" {
				result = fmt.Sprintf("WHERE %%s", result)
			}
			return result, nil
		}

	`,
		name,
		capitalName,
	)

	code := strings.ReplaceAll(template, "__BACKTICK__", "`")
	filePath := fmt.Sprintf("repositories/%s_repository/%s_repository_map.go", name, name)

	save(filePath, code)
}
