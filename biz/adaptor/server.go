package adaptor

import "github.com/xh-polaris/psych-model/biz/adaptor/controller"

type Server struct {
	controller.IUnitAppConfigController
	controller.IAppController
}
