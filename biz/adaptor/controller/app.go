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
	AppGet(ctx context.Context, req *m.AppGetReq) (res *m.AppGetResp, err error)
	AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error)
}

type AppController struct {
	UnitAppConfigService *service.UnitAppConfigService
}

var AppControllerSet = wire.NewSet(
	wire.Struct(new(AppController), "*"),
	wire.Bind(new(IAppController), new(*AppController)),
)

func (a AppController) AppCreate(ctx context.Context, req *m.AppCreateReq) (res *m.AppCreateResp, err error) {
	return a.AppCreate(ctx, req)
}

func (a AppController) AppUpdate(ctx context.Context, req *m.AppUpdateReq) (res *basic.Response, err error) {
	return a.AppUpdate(ctx, req)
}

func (a AppController) AppGet(ctx context.Context, req *m.AppGetReq) (res *m.AppGetResp, err error) {
	return a.AppGet(ctx, req)
}

func (a AppController) AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error) {
	return a.AppDelete(ctx, req)
}
