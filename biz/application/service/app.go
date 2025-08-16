package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	m "github.com/xh-polaris/psych-idl/kitex_gen/model"
	"github.com/xh-polaris/psych-model/biz/infrastructure/consts"
	appmapper "github.com/xh-polaris/psych-model/biz/infrastructure/mapper/app"
	mdlmapper "github.com/xh-polaris/psych-model/biz/infrastructure/mapper/model"
	"github.com/xh-polaris/psych-model/biz/infrastructure/util/result"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAppService interface {
	AppCreate(ctx context.Context, req *m.AppCreateReq) (res *m.AppCreateResp, err error)
	AppUpdate(ctx context.Context, req *m.AppUpdateReq) (res *basic.Response, err error)
	AppGetByConfigId(ctx context.Context, req *m.AppGetByConfigIdReq) (res *m.AppGetByConfigIdResp, err error)
	AppGetById(ctx context.Context, req *m.AppGetByIdReq) (res *m.AppGetByIdResp, err error)
	AppList(ctx context.Context, req *m.AppListReq) (res *m.AppListResp, err error)
	AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error)
}

type AppService struct {
	AppMapper   *appmapper.MongoMapper
	ModelMapper *mdlmapper.MongoMapper
}

var AppServiceSet = wire.NewSet(
	wire.Struct(new(AppService), "*"),
	wire.Bind(new(IAppService), new(*AppService)),
)

func (a *AppService) AppCreate(ctx context.Context, req *m.AppCreateReq) (res *m.AppCreateResp, err error) {
	var app *m.AppData
	switch appData := req.App.App.(type) {
	case *m.AppData_ChatApp:
		app, err = a.appCreate(ctx, &m.AppData{App: appData}, req.ConfigId, consts.ChatApp)
	case *m.AppData_TtsApp:
		app, err = a.appCreate(ctx, &m.AppData{App: appData}, req.ConfigId, consts.TtsApp)
	case *m.AppData_AsrApp:
		app, err = a.appCreate(ctx, &m.AppData{App: appData}, req.ConfigId, consts.AsrApp)
	case *m.AppData_ReportApp:
		app, err = a.appCreate(ctx, &m.AppData{App: appData}, req.ConfigId, consts.ReportApp)
	default:
		return nil, consts.ErrInvalidParams
	}
	if err != nil {
		return nil, err
	}
	return &m.AppCreateResp{App: app}, nil
}

func (a *AppService) appCreate(ctx context.Context, app *m.AppData, configId string, types int32) (res *m.AppData, err error) {
	appWrap := appGen2DB(app)
	appWrap.ConfigId = configId

	// 插入app db
	id, err := a.AppMapper.InsertWithEcho(ctx, appWrap)
	if err != nil {
		return nil, err
	}

	// 更新model db
	if err = a.ModelMapper.UpdateAppId(ctx, configId, types, id); err != nil {
		return nil, err
	}

	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	appWrap.ID = hex
	return appDB2Gen(appWrap, true), nil
}

func (a *AppService) AppUpdate(ctx context.Context, req *m.AppUpdateReq) (res *basic.Response, err error) {
	switch appData := req.App.App.(type) {
	case *m.AppData_ChatApp, *m.AppData_TtsApp, *m.AppData_AsrApp, *m.AppData_ReportApp:
		appWrap := appGen2DB(&m.AppData{App: appData})
		if appWrap == nil {
			return nil, consts.ErrInvalidParams
		}
		err = a.AppMapper.Update(ctx, appWrap)
	default:
		return nil, consts.ErrInvalidParams
	}
	return result.ResponseOk(), nil
}

func (a *AppService) AppGetByConfigId(ctx context.Context, req *m.AppGetByConfigIdReq) (res *m.AppGetByConfigIdResp, err error) {
	configId := req.ConfigId
	// 查询 app
	apps, err := a.AppMapper.FindBatchByConfigId(ctx, configId)
	if err != nil {
		return nil, err
	}

	resApps := make([]*m.AppData, len(apps))
	for idx, app := range apps {
		gen := appDB2Gen(app, true)
		resApps[idx] = gen
	}

	return &m.AppGetByConfigIdResp{
		Apps: resApps,
	}, nil
}

func (a *AppService) AppGetById(ctx context.Context, req *m.AppGetByIdReq) (res *m.AppGetByIdResp, err error) {
	// 查询 app
	app, err := a.AppMapper.FindOneById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &m.AppGetByIdResp{App: appDB2Gen(app, true)}, nil
}

func (a *AppService) AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error) {
	// 删除unit_model对应字段，并获取appId
	configId := req.ConfigId
	appId, err := a.ModelMapper.DeleteAppId(ctx, configId, req.Type)
	if err != nil {
		return nil, err
	}
	// 删除app
	if err = a.AppMapper.DeleteOneById(ctx, appId); err != nil {
		return nil, err
	}

	return result.ResponseOk(), nil
}

func (a *AppService) AppList(ctx context.Context, req *m.AppListReq) (res *m.AppListResp, err error) {
	p := req.GetPaginationOptions()
	data, total, err := a.AppMapper.List(ctx, p, req.Type)
	if err != nil {
		return nil, err
	}

	var apps []*m.AppData
	for _, app := range data {
		apps = append(apps, appDB2Gen(&app, true))
	}

	return &m.AppListResp{
		Apps:     apps,
		Total:    total,
		Page:     *p.Page,
		PageSize: *p.Limit,
	}, nil
}

func appDB2Gen(app *appmapper.AppWrap, admin bool) *m.AppData {
	base := app.App.GetBase()
	var appBase *m.App
	if admin {
		appBase = &m.App{
			Id:          app.ID.Hex(),
			Name:        base.Name,
			Description: base.Description,
			Lang:        base.Lang,
			Platform:    base.Platform,
			Provider:    base.Provider,
			Url:         base.Url,
			AppId:       base.AppId,
			AccessKey:   base.AccessKey,
			Stream:      base.Stream,
			Level:       base.Level,
			Status:      base.Status,
			CreateTime:  base.CreateTime.Unix(),
			ExpireTime:  base.ExpireTime,
			UpdateTime:  base.UpdateTime.Unix(),
		}
	} else {
		appBase = &m.App{
			Id:          app.ID.Hex(),
			Name:        base.Name,
			Description: base.Description,
			Lang:        base.Lang,
			Platform:    base.Platform,
			Provider:    base.Provider,
			Url:         base.Url,
			Stream:      base.Stream,
			Level:       base.Level,
			Status:      base.Status,
			CreateTime:  base.CreateTime.Unix(),
			ExpireTime:  base.ExpireTime,
			UpdateTime:  base.UpdateTime.Unix(),
		}
	}

	switch app.Type {
	case consts.ChatApp:
		return &m.AppData{
			Type: consts.ChatApp,
			App: &m.AppData_ChatApp{
				ChatApp: &m.ChatApp{
					App: appBase,
				},
			},
		}
	case consts.TtsApp:
		a := app.App.(*appmapper.TtsApp)
		return &m.AppData{
			Type: consts.TtsApp,
			App: &m.AppData_TtsApp{TtsApp: &m.TtsApp{
				App:        appBase,
				Namespace:  a.Namespace,
				Speaker:    a.Speaker,
				ResourceId: a.ResourceId,
				AudioParams: &m.TtsApp_AudioParam{
					Format:       a.AudioParam.Format,
					Codec:        a.AudioParam.Codec,
					Rate:         a.AudioParam.Rate,
					Bits:         a.AudioParam.Bits,
					Channels:     a.AudioParam.Channels,
					SpeechRate:   a.AudioParam.SpeechRate,
					LoudnessRate: a.AudioParam.LoudnessRate,
					Lang:         a.AudioParam.Lang,
					ResultType:   a.AudioParam.ResultType,
				},
			}},
		}
	case consts.AsrApp:
		a := app.App.(*appmapper.AsrApp)
		return &m.AppData{
			Type: consts.AsrApp,
			App: &m.AppData_AsrApp{AsrApp: &m.AsrApp{
				App:        appBase,
				Format:     a.Format,
				Codec:      a.Codec,
				Rate:       a.Rate,
				Bits:       a.Bits,
				Channels:   a.Channels,
				ModelName:  a.ModelName,
				EnablePunc: a.EnablePunc,
				EnableDdc:  a.EnableDdc,
				ResultType: a.ResultType,
				ResourceId: a.ResourceId,
			}},
		}
	case consts.ReportApp:
		return &m.AppData{
			Type: consts.ReportApp,
			App:  &m.AppData_ReportApp{ReportApp: &m.ReportApp{App: appBase}},
		}
	default:
		return nil // 或者 panic("unsupported app type")
	}
}

func appGen2DB(appData *m.AppData) *appmapper.AppWrap {
	app := appData.App
	switch a := app.(type) {
	case *m.AppData_ChatApp:
		baseApp := a.ChatApp.App
		var oid primitive.ObjectID
		var err error
		if baseApp.Id != "" {
			oid, err = primitive.ObjectIDFromHex(baseApp.Id)
			if err != nil {
				return nil
			}
		}
		return &appmapper.AppWrap{
			ID:   oid,
			Type: consts.ChatApp,
			App: &appmapper.ChatApp{
				AppBase: appmapper.AppBase{
					Name:        baseApp.Name,
					Description: baseApp.Description,
					Lang:        baseApp.Lang,
					Platform:    baseApp.Platform,
					Provider:    baseApp.Provider,
					Url:         baseApp.Url,
					AppId:       baseApp.AppId,
					AccessKey:   baseApp.AccessKey,
					Stream:      baseApp.Stream,
					Level:       baseApp.Level,
					Status:      baseApp.Status,
					ExpireTime:  baseApp.ExpireTime,
				},
			},
		}

	case *m.AppData_TtsApp:
		ttsApp := a.TtsApp
		baseApp := ttsApp.App
		var oid primitive.ObjectID
		var err error
		if baseApp.Id != "" {
			oid, err = primitive.ObjectIDFromHex(baseApp.Id)
			if err != nil {
				return nil
			}
		}
		return &appmapper.AppWrap{
			ID:   oid,
			Type: consts.TtsApp,
			App: &appmapper.TtsApp{
				AppBase: appmapper.AppBase{
					Name:        baseApp.Name,
					Description: baseApp.Description,
					Lang:        baseApp.Lang,
					Platform:    baseApp.Platform,
					Provider:    baseApp.Provider,
					Url:         baseApp.Url,
					AppId:       baseApp.AppId,
					AccessKey:   baseApp.AccessKey,
					Stream:      baseApp.Stream,
					Level:       baseApp.Level,
					Status:      baseApp.Status,
					ExpireTime:  baseApp.ExpireTime,
				},
				Namespace:  ttsApp.Namespace,
				Speaker:    ttsApp.Speaker,
				ResourceId: ttsApp.ResourceId,
				AudioParam: appmapper.AudioParam{
					Format:       ttsApp.AudioParams.Format,
					Codec:        ttsApp.AudioParams.Codec,
					Rate:         ttsApp.AudioParams.Rate,
					Bits:         ttsApp.AudioParams.Bits,
					Channels:     ttsApp.AudioParams.Channels,
					SpeechRate:   ttsApp.AudioParams.SpeechRate,
					LoudnessRate: ttsApp.AudioParams.LoudnessRate,
					Lang:         ttsApp.AudioParams.Lang,
					ResultType:   ttsApp.AudioParams.ResultType,
				},
			},
		}
	case *m.AppData_AsrApp:
		asrApp := a.AsrApp
		baseApp := asrApp.App
		var oid primitive.ObjectID
		var err error
		if baseApp.Id != "" {
			oid, err = primitive.ObjectIDFromHex(baseApp.Id)
			if err != nil {
				return nil
			}
		}
		return &appmapper.AppWrap{
			ID:   oid,
			Type: consts.AsrApp,
			App: &appmapper.AsrApp{
				AppBase: appmapper.AppBase{
					Name:        baseApp.Name,
					Description: baseApp.Description,
					Lang:        baseApp.Lang,
					Platform:    baseApp.Platform,
					Provider:    baseApp.Provider,
					Url:         baseApp.Url,
					AppId:       baseApp.AppId,
					AccessKey:   baseApp.AccessKey,
					Stream:      baseApp.Stream,
					Level:       baseApp.Level,
					Status:      baseApp.Status,
					ExpireTime:  baseApp.ExpireTime,
				},
				Format:     asrApp.Format,
				Codec:      asrApp.Codec,
				Rate:       asrApp.Rate,
				Bits:       asrApp.Bits,
				ResourceId: asrApp.ResourceId,
				Channels:   asrApp.Channels,
				ModelName:  asrApp.ModelName,
				EnablePunc: asrApp.EnablePunc,
				EnableDdc:  asrApp.EnableDdc,
				ResultType: asrApp.ResultType,
			},
		}
	case *m.AppData_ReportApp:
		reportApp := a.ReportApp
		baseApp := reportApp.App
		var oid primitive.ObjectID
		var err error
		if baseApp.Id != "" {
			oid, err = primitive.ObjectIDFromHex(baseApp.Id)
			if err != nil {
				return nil
			}
		}
		return &appmapper.AppWrap{
			ID:   oid,
			Type: consts.ReportApp,
			App: &appmapper.ReportApp{
				AppBase: appmapper.AppBase{
					Name:        baseApp.Name,
					Description: baseApp.Description,
					Lang:        baseApp.Lang,
					Platform:    baseApp.Platform,
					Provider:    baseApp.Provider,
					Url:         baseApp.Url,
					AppId:       baseApp.AppId,
					AccessKey:   baseApp.AccessKey,
					Stream:      baseApp.Stream,
					Level:       baseApp.Level,
					Status:      baseApp.Status,
					ExpireTime:  baseApp.ExpireTime,
				},
			},
		}
	default:
		return nil // panic("unsupported app type")
	}
}
