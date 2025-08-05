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
	AppGetByConfigId(ctx context.Context, req *m.AppGetByConfigIdReq) (res *m.AppGetByConfigIdResp, err error)
	AppGetById(ctx context.Context, req *m.AppGetByIdReq) (res *m.AppGetByIdResp, err error)
	AppList(ctx context.Context, req *m.AppListReq) (res *m.AppListResp, err error)
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

func (a *AppController) AppGetByConfigId(ctx context.Context, req *m.AppGetByConfigIdReq) (res *m.AppGetByConfigIdResp, err error) {
	return a.AppService.AppGetByConfigId(ctx, req)
}

func (a *AppController) AppGetById(ctx context.Context, req *m.AppGetByIdReq) (res *m.AppGetByIdResp, err error) {
	return a.AppService.AppGetById(ctx, req)
}

func (a *AppController) AppList(ctx context.Context, req *m.AppListReq) (res *m.AppListResp, err error) {
	return a.AppService.AppList(ctx, req)
}

func (a *AppController) AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error) {
	return a.AppService.AppDelete(ctx, req)
}
