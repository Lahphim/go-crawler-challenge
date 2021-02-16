package helpers

import (
	"regexp"

	"github.com/beego/beego/v2/server/web"
	"github.com/iancoleman/strcase"
)

type NestPreparer interface {
	NestPrepare()
}

// SetControllerAttributes sets some attributes for controller
func SetControllerAttributes(controller *web.Controller) {
	controllerName, actionName := controller.GetControllerAndAction()

	re := regexp.MustCompile(`Controller$`)
	controllerName = re.ReplaceAllString(controllerName, "")

	controller.Data["ControllerName"] = strcase.ToKebab(controllerName)
	controller.Data["ActionName"] = strcase.ToKebab(actionName)
}

// SetFlashMessageLayout sets flash message layout for controller
func SetFlashMessageLayout(controller *web.Controller) {
	controller.LayoutSections = make(map[string]string)
	controller.LayoutSections["FlashMessage"] = "shared/alert.html"

	app, ok := controller.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}
}
