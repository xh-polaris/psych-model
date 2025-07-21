package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/psych-model/biz/adaptor/controller"
	"github.com/xh-polaris/psych-model/biz/application/service"
	"github.com/xh-polaris/psych-model/biz/infrastructure/config"
	"github.com/xh-polaris/psych-model/biz/infrastructure/mapper/app"
	"github.com/xh-polaris/psych-model/biz/infrastructure/mapper/model"
)

var ApplicationSet = wire.NewSet(
	service.AppServiceSet,
	service.UnitAppConfigServiceSet,
)

var MapperSet = wire.NewSet(
	model.NewMongoMapper,
	app.NewMongoMapper,
)

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	MapperSet,
)

var ControllerSet = wire.NewSet(
	controller.AppControllerSet,
	controller.UnitAppConfigControllerSet,
)

var ServerProvider = wire.NewSet(
	ControllerSet,
	ApplicationSet,
	InfrastructureSet,
)
