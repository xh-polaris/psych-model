package controller

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	m "github.com/xh-polaris/psych-idl/kitex_gen/model"
	"github.com/xh-polaris/psych-model/biz/application/service"
)

type IUnitAppConfigController interface {
	UnitAppConfigCreate(ctx context.Context, req *m.UnitAppConfigCreateReq) (res *m.UnitAppConfigCreateResp, err error)
	UnitAppConfigUpdate(ctx context.Context, req *m.UnitAppConfigUpdateReq) (res *basic.Response, err error)
	UnitAppConfigGet(ctx context.Context, req *m.UnitAppConfigGetReq) (res *m.UnitAppConfigGetResp, err error)
}

type UnitAppConfigController struct {
	UnitAppConfigService *service.UnitAppConfigService
}

var UnitAppConfigControllerSet = wire.NewSet(
	wire.Struct(new(UnitAppConfigController), "*"),
	wire.Bind(new(IUnitAppConfigController), new(*UnitAppConfigController)),
)

func (u UnitAppConfigController) UnitAppConfigCreate(ctx context.Context, req *m.UnitAppConfigCreateReq) (res *m.UnitAppConfigCreateResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (u UnitAppConfigController) UnitAppConfigUpdate(ctx context.Context, req *m.UnitAppConfigUpdateReq) (res *basic.Response, err error) {
	//TODO implement me
	panic("implement me")
}

func (u UnitAppConfigController) UnitAppConfigGet(ctx context.Context, req *m.UnitAppConfigGetReq) (res *m.UnitAppConfigGetResp, err error) {
	//TODO implement me
	panic("implement me")
}
