package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateHandlerInterface(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package %s_handler

		import (
			"net/http"

			"github.com/yunarsuanto/base-go/infra/initiator/service"
		)

		type %sHandlerInterface interface {
			List%s(w http.ResponseWriter, r *http.Request)
			Create%s(w http.ResponseWriter, r *http.Request)
			Update%s(w http.ResponseWriter, r *http.Request)
			Delete%s(w http.ResponseWriter, r *http.Request)
		}

		func New%sHandler(serviceCtx *service.ServiceCtx) %sHandlerInterface {
			return &handler{
				serviceCtx,
			}
		}


	`,
		name,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
	)

	code := strings.ReplaceAll(template, "__BACKTICK__", "`")
	filePath := fmt.Sprintf("handler/%s_handler/%s_handler_interface.go", name, name)

	save(filePath, code)
}
