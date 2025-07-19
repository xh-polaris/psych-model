package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	m "github.com/xh-polaris/psych-idl/kitex_gen/model"
	untmapper "github.com/xh-polaris/psych-model/biz/infrastructure/mapper/unit"
)

type IUnitAppConfigService interface {
	UnitAppConfigCreate(ctx context.Context, req *m.UnitAppConfigCreateReq) (res *m.UnitAppConfigCreateResp, err error)
	UnitAppConfigUpdate(ctx context.Context, req *m.UnitAppConfigUpdateReq) (res *basic.Response, err error)
	UnitAppConfigGet(ctx context.Context, req *m.UnitAppConfigGetReq) (res *m.UnitAppConfigGetResp, err error)
}

type UnitAppConfigService struct {
	UnitAppConfigMapper *untmapper.MongoMapper
}

var UnitAppConfigServiceSet = wire.NewSet(
	wire.Struct(new(UnitAppConfigService), "*"),
	wire.Bind(new(IUnitAppConfigService), new(*UnitAppConfigService)),
)

func (u UnitAppConfigService) UnitAppConfigCreate(ctx context.Context, req *m.UnitAppConfigCreateReq) (res *m.UnitAppConfigCreateResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (u UnitAppConfigService) UnitAppConfigUpdate(ctx context.Context, req *m.UnitAppConfigUpdateReq) (res *basic.Response, err error) {
	//TODO implement me
	panic("implement me")
}

func (u UnitAppConfigService) UnitAppConfigGet(ctx context.Context, req *m.UnitAppConfigGetReq) (res *m.UnitAppConfigGetResp, err error) {
	//TODO implement me
	panic("implement me")
}
