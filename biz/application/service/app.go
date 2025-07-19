package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	m "github.com/xh-polaris/psych-idl/kitex_gen/model"
	appmapper "github.com/xh-polaris/psych-model/biz/infrastructure/mapper/app"
)

type IAppService interface {
	AppCreate(ctx context.Context, req *m.AppCreateReq) (res *m.AppCreateResp, err error)
	AppUpdate(ctx context.Context, req *m.AppUpdateReq) (res *basic.Response, err error)
	AppGet(ctx context.Context, req *m.AppGetReq) (res *m.AppGetResp, err error)
	AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error)
}

type AppService struct {
	AppMapper *appmapper.MongoMapper
}

var AppServiceSet = wire.NewSet(
	wire.Struct(new(AppService), "*"),
	wire.Bind(new(IAppService), new(*AppService)),
)

func (a AppService) AppCreate(ctx context.Context, req *m.AppCreateReq) (res *m.AppCreateResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AppService) AppUpdate(ctx context.Context, req *m.AppUpdateReq) (res *basic.Response, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AppService) AppGet(ctx context.Context, req *m.AppGetReq) (res *m.AppGetResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AppService) AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error) {
	//TODO implement me
	panic("implement me")
}
