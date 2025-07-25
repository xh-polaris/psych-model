package controller

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	m "github.com/xh-polaris/psych-idl/kitex_gen/model"
	"github.com/xh-polaris/psych-model/biz/application/service"
)

type IAppController interface {
	AppCreate(ctx context.Context, req *m.AppCreateReq) (res *m.AppCreateResp, err error)
	AppUpdate(ctx context.Context, req *m.AppUpdateReq) (res *basic.Response, err error)
	AppGetByUnitIdReq(ctx context.Context, req *m.AppGetByUnitIdReq) (res *m.AppGetByUnitIdResp, err error)
	AppGetPagesReq(ctx context.Context, req *m.AppGetPagesReq) (res *m.AppGetPagesResp, err error)
	AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error)
}

type AppController struct {
	AppService *service.AppService
}

var AppControllerSet = wire.NewSet(
	wire.Struct(new(AppController), "*"),
	wire.Bind(new(IAppController), new(*AppController)),
)

func (a *AppController) AppCreate(ctx context.Context, req *m.AppCreateReq) (res *m.AppCreateResp, err error) {
	return a.AppService.AppCreate(ctx, req)
}

func (a *AppController) AppUpdate(ctx context.Context, req *m.AppUpdateReq) (res *basic.Response, err error) {
	return a.AppService.AppUpdate(ctx, req)
}

func (a *AppController) AppGetByUnitIdReq(ctx context.Context, req *m.AppGetByUnitIdReq) (res *m.AppGetByUnitIdResp, err error) {
	return a.AppService.AppGetByUnitIdReq(ctx, req)
}

func (a *AppController) AppGetPagesReq(ctx context.Context, req *m.AppGetPagesReq) (res *m.AppGetPagesResp, err error) {
	return a.AppService.AppGetPagesReq(ctx, req)
}

func (a *AppController) AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error) {
	return a.AppService.AppDelete(ctx, req)
}
