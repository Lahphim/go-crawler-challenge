package helpers

import (
	"regexp"

	"github.com/beego/beego/v2/server/web"
	"github.com/iancoleman/strcase"
)

// SetControllerAttributes sets some attributes for controller
func SetControllerAttributes(controller *web.Controller) {
	controllerName, actionName := controller.GetControllerAndAction()

	re := regexp.MustCompile(`Controller$`)
	controllerName = re.ReplaceAllString(controllerName, "")

	controller.Data["ControllerName"] = strcase.ToKebab(controllerName)
	controller.Data["ActionName"] = strcase.ToKebab(actionName)
}
