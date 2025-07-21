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
	"time"
)

type IAppService interface {
	AppCreate(ctx context.Context, req *m.AppCreateReq) (res *m.AppCreateResp, err error)
	AppUpdate(ctx context.Context, req *m.AppUpdateReq) (res *basic.Response, err error)
	AppGet(ctx context.Context, req *m.AppGetReq) (res *m.AppGetResp, err error)
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
	switch detail := req.AppDetail.(type) {
	case *m.AppCreateReq_ChatApp:
		// detail 是 *AppCreateReq_ChatApp 类型
		res, err := a.createChat(ctx, detail, req.UnitAppConfig)
		if err != nil {
			return nil, err
		}
		return &m.AppCreateResp{
			AppDetail: res,
		}, err

	case *m.AppCreateReq_TtsApp:
		// detail 是 *AppCreateReq_TtsApp 类型
		res, err := a.createTts(ctx, detail, req.UnitAppConfig)
		if err != nil {
			return nil, err
		}
		return &m.AppCreateResp{
			AppDetail: res,
		}, err

	case *m.AppCreateReq_AsrApp:
		// detail 是 *AppCreateReq_AsrApp 类型
		res, err := a.createAsr(ctx, detail, req.UnitAppConfig)
		if err != nil {
			return nil, err
		}
		return &m.AppCreateResp{
			AppDetail: res,
		}, err

	case *m.AppCreateReq_ReportApp:
		// detail 是 *AppCreateReq_ReportApp 类型
		res, err := a.createReport(ctx, detail, req.UnitAppConfig)
		if err != nil {
			return nil, err
		}
		return &m.AppCreateResp{
			AppDetail: res,
		}, err
	default:
		return nil, consts.ErrInvalidParams
	}
}

func (a *AppService) AppUpdate(ctx context.Context, req *m.AppUpdateReq) (res *basic.Response, err error) {
	switch detail := req.AppDetail.(type) {
	case *m.AppUpdateReq_ChatApp:
		// detail 是 *AppCreateReq_ChatApp 类型
		app := detail.ChatApp.App
		oid, err := primitive.ObjectIDFromHex(app.Id)
		if err != nil {
			return nil, err
		}
		err = a.AppMapper.Update(ctx, &appmapper.ChatApp{
			App: appmapper.AppBase{
				ID:          oid,
				Name:        app.Name,
				Description: app.Description,
				Lang:        app.Lang,
				Platform:    app.Platform,
				Url:         app.Url,
				Appid:       app.Appid,
				Auth:        app.Auth,
				Stream:      app.Stream,
				Level:       app.Level,
				Status:      app.Status,
				ExpireTime:  app.ExpireTime,
			},
		})

	case *m.AppUpdateReq_TtsApp:
		// detail 是 *AppCreateReq_TtsApp 类型
		tts := detail.TtsApp
		app := tts.App
		oid, err := primitive.ObjectIDFromHex(app.Id)
		if err != nil {
			return nil, err
		}
		err = a.AppMapper.Update(ctx, &appmapper.TtsApp{
			App: appmapper.AppBase{
				ID:          oid,
				Name:        app.Name,
				Description: app.Description,
				Lang:        app.Lang,
				Platform:    app.Platform,
				Url:         app.Url,
				Appid:       app.Appid,
				Auth:        app.Auth,
				Stream:      app.Stream,
				Level:       app.Level,
				Status:      app.Status,
				ExpireTime:  app.ExpireTime,
			},
			Namespace:  tts.Namespace,
			Speaker:    tts.Speaker,
			ResourceId: tts.ResourceId,
			AudioParam: appmapper.AudioParam{
				Format:       tts.AudioParams.Format,
				Rate:         tts.AudioParams.Rate,
				Bit:          tts.AudioParams.Bit,
				SpeechRate:   tts.AudioParams.SpeechRate,
				LoudnessRate: tts.AudioParams.LoudnessRate,
				Lang:         tts.AudioParams.Lang,
			},
		})

	case *m.AppUpdateReq_AsrApp:
		// detail 是 *AppCreateReq_AsrApp 类型
		asr := detail.AsrApp
		app := asr.App
		oid, err := primitive.ObjectIDFromHex(app.Id)
		if err != nil {
			return nil, err
		}
		err = a.AppMapper.Update(ctx, &appmapper.AsrApp{
			App: appmapper.AppBase{
				ID:          oid,
				Name:        app.Name,
				Description: app.Description,
				Lang:        app.Lang,
				Platform:    app.Platform,
				Url:         app.Url,
				Appid:       app.Appid,
				Auth:        app.Auth,
				Stream:      app.Stream,
				Level:       app.Level,
				Status:      app.Status,
				ExpireTime:  app.ExpireTime,
			},
			Format:     asr.Format,
			Codec:      asr.Codec,
			Rate:       asr.Rate,
			Bits:       asr.Bits,
			Channels:   asr.Channels,
			ModelName:  asr.ModelName,
			EnablePunc: asr.EnablePunc,
			EnableDdc:  asr.EnableDdc,
			ResultType: asr.ResultType,
		})

	case *m.AppUpdateReq_ReportApp:
		// detail 是 *AppCreateReq_ReportApp 类型
		report := detail.ReportApp
		app := report.App
		oid, err := primitive.ObjectIDFromHex(app.Id)
		if err != nil {
			return nil, err
		}
		err = a.AppMapper.Update(ctx, &appmapper.ReportApp{
			App: appmapper.AppBase{
				ID:          oid,
				Name:        app.Name,
				Description: app.Description,
				Lang:        app.Lang,
				Platform:    app.Platform,
				Url:         app.Url,
				Appid:       app.Appid,
				Auth:        app.Auth,
				Stream:      app.Stream,
				Level:       app.Level,
				Status:      app.Status,
				ExpireTime:  app.ExpireTime,
			},
		})
	default:
		return nil, consts.ErrInvalidParams
	}
	return result.ResponseOk(), nil
}

func (a *AppService) AppGet(ctx context.Context, req *m.AppGetReq) (res *m.AppGetResp, err error) {
	unitId := req.UnitAppConfig.UnitId
	app, err := a.AppMapper.FindByUnitId(ctx, unitId, req.Type)
	if err != nil {
		return nil, err
	}
	base := app.GetBase()
	resApp := &m.App{
		Id:          base.ID.Hex(),
		Name:        base.Name,
		Description: base.Description,
		Lang:        base.Lang,
		Platform:    base.Platform,
		Url:         base.Url,
		Appid:       base.Appid,
		Auth:        base.Auth,
		Stream:      base.Stream,
		Level:       base.Level,
		Status:      base.Status,
		CreateTime:  base.CreateTime.Unix(),
		ExpireTime:  base.ExpireTime,
		UpdateTime:  base.UpdateTime.Unix(),
	}
	switch detail := app.(type) {
	case *appmapper.ChatApp:
		res.ChatApp = &m.ChatApp{
			App: resApp,
		}
	case *appmapper.TtsApp:
		res.TtsApp = &m.TtsApp{
			App:        resApp,
			Namespace:  detail.Namespace,
			Speaker:    detail.Speaker,
			ResourceId: detail.ResourceId,
			AudioParams: &m.TtsApp_AudioParam{
				Format:       detail.AudioParam.Format,
				Rate:         detail.AudioParam.Rate,
				Bit:          detail.AudioParam.Bit,
				SpeechRate:   detail.AudioParam.SpeechRate,
				LoudnessRate: detail.AudioParam.LoudnessRate,
				Lang:         detail.AudioParam.Lang,
			},
		}
	case *appmapper.AsrApp:
		res.AsrApp = &m.AsrApp{
			App:        resApp,
			Format:     detail.Format,
			Codec:      detail.Codec,
			Rate:       detail.Rate,
			Bits:       detail.Bits,
			Channels:   detail.Channels,
			ModelName:  detail.ModelName,
			EnablePunc: detail.EnablePunc,
			EnableDdc:  detail.EnableDdc,
			ResultType: detail.ResultType,
		}
	case *appmapper.ReportApp:
		res.ReportApp = &m.ReportApp{
			App: resApp,
		}
	default:
		return nil, consts.ErrInvalidParams
	}
	return res, nil
}

func (a *AppService) AppDelete(ctx context.Context, req *m.AppDeleteReq) (res *basic.Response, err error) {
	// 删除app
	unitId := req.UnitAppConfig.UnitId
	if err = a.AppMapper.Delete(ctx, unitId, req.Type); err != nil {
		return nil, err
	}

	// 删除unit_model对应字段
	if err = a.ModelMapper.UpdateAppid(ctx, unitId, req.Type, ""); err != nil {
		return nil, err
	}

	return result.ResponseOk(), nil
}

func (a *AppService) createChat(ctx context.Context, detail *m.AppCreateReq_ChatApp, model *m.UnitAppConfig) (*m.AppCreateResp_ChatApp, error) {
	now := time.Now()

	// 插入app表
	chatApp := detail.ChatApp
	app := chatApp.App
	id, err := a.AppMapper.InsertWithEcho(ctx, &appmapper.ChatApp{
		App: appmapper.AppBase{
			Name:        app.Name,
			Description: app.Description,
			Lang:        app.Lang,
			Platform:    app.Platform,
			Url:         app.Url,
			Appid:       app.Appid,
			Auth:        app.Auth,
			Stream:      app.Stream,
			Level:       app.Level,
			Status:      app.Status,
			ExpireTime:  app.ExpireTime,
			CreateTime:  now,
			UpdateTime:  now,
			Type:        consts.ChatApp,
		},
	})
	if err != nil {
		return nil, err
	}

	// 插入unit_model表
	err = a.ModelMapper.UpdateAppid(ctx, model.Id, consts.ChatApp, id)
	if err != nil {
		return nil, err
	}

	app.Id = id
	app.CreateTime = now.Unix()
	app.UpdateTime = now.Unix()
	return &m.AppCreateResp_ChatApp{
		ChatApp: &m.ChatApp{
			App: app,
		},
	}, nil
}

func (a *AppService) createTts(ctx context.Context, detail *m.AppCreateReq_TtsApp, model *m.UnitAppConfig) (*m.AppCreateResp_TtsApp, error) {
	now := time.Now()

	// 插入app表
	ttsApp := detail.TtsApp
	app := ttsApp.App
	id, err := a.AppMapper.InsertWithEcho(ctx, &appmapper.TtsApp{
		App: appmapper.AppBase{
			Name:        app.Name,
			Description: app.Description,
			Lang:        app.Lang,
			Platform:    app.Platform,
			Url:         app.Url,
			Appid:       app.Appid,
			Auth:        app.Auth,
			Stream:      app.Stream,
			Level:       app.Level,
			Status:      app.Status,
			ExpireTime:  app.ExpireTime,
			CreateTime:  now,
			UpdateTime:  now,
			Type:        consts.TtsApp,
		},
		Namespace:  ttsApp.Namespace,
		Speaker:    ttsApp.Speaker,
		ResourceId: ttsApp.ResourceId,
		AudioParam: appmapper.AudioParam{
			Format:       ttsApp.AudioParams.Format,
			Rate:         ttsApp.AudioParams.Rate,
			Bit:          ttsApp.AudioParams.Bit,
			SpeechRate:   ttsApp.AudioParams.SpeechRate,
			LoudnessRate: ttsApp.AudioParams.LoudnessRate,
			Lang:         ttsApp.AudioParams.Lang,
		},
	})
	if err != nil {
		return nil, err
	}

	// 插入unit_model表
	err = a.ModelMapper.UpdateAppid(ctx, model.Id, consts.TtsApp, id)
	if err != nil {
		return nil, err
	}

	app.Id = id
	app.CreateTime = now.Unix()
	app.UpdateTime = now.Unix()
	return &m.AppCreateResp_TtsApp{
		TtsApp: &m.TtsApp{
			App:        app,
			Namespace:  ttsApp.Namespace,
			Speaker:    ttsApp.Speaker,
			ResourceId: ttsApp.ResourceId,
			AudioParams: &m.TtsApp_AudioParam{
				Format:       ttsApp.AudioParams.Format,
				Rate:         ttsApp.AudioParams.Rate,
				Bit:          ttsApp.AudioParams.Bit,
				SpeechRate:   ttsApp.AudioParams.SpeechRate,
				LoudnessRate: ttsApp.AudioParams.LoudnessRate,
				Lang:         ttsApp.AudioParams.Lang,
			},
		},
	}, nil
}

func (a *AppService) createAsr(ctx context.Context, detail *m.AppCreateReq_AsrApp, model *m.UnitAppConfig) (*m.AppCreateResp_AsrApp, error) {
	now := time.Now()

	// 插入app表
	asrApp := detail.AsrApp
	app := asrApp.App
	id, err := a.AppMapper.InsertWithEcho(ctx, &appmapper.AsrApp{
		App: appmapper.AppBase{
			Name:        app.Name,
			Description: app.Description,
			Lang:        app.Lang,
			Platform:    app.Platform,
			Url:         app.Url,
			Appid:       app.Appid,
			Auth:        app.Auth,
			Stream:      app.Stream,
			Level:       app.Level,
			Status:      app.Status,
			ExpireTime:  app.ExpireTime,
			CreateTime:  now,
			UpdateTime:  now,
			Type:        consts.AsrApp,
		},
		Format:     asrApp.Format,
		Codec:      asrApp.Codec,
		Rate:       asrApp.Rate,
		Bits:       asrApp.Bits,
		Channels:   asrApp.Channels,
		ModelName:  asrApp.ModelName,
		EnablePunc: asrApp.EnablePunc,
		EnableDdc:  asrApp.EnableDdc,
		ResultType: asrApp.ResultType,
	})
	if err != nil {
		return nil, err
	}

	// 插入unit_model表
	err = a.ModelMapper.UpdateAppid(ctx, model.Id, consts.AsrApp, id)
	if err != nil {
		return nil, err
	}

	app.Id = id
	app.CreateTime = now.Unix()
	app.UpdateTime = now.Unix()
	return &m.AppCreateResp_AsrApp{
		AsrApp: &m.AsrApp{
			App:        app,
			Format:     asrApp.Format,
			Codec:      asrApp.Codec,
			Rate:       asrApp.Rate,
			Bits:       asrApp.Bits,
			Channels:   asrApp.Channels,
			ModelName:  asrApp.ModelName,
			EnablePunc: asrApp.EnablePunc,
			EnableDdc:  asrApp.EnableDdc,
			ResultType: asrApp.ResultType,
		},
	}, nil
}

func (a *AppService) createReport(ctx context.Context, detail *m.AppCreateReq_ReportApp, model *m.UnitAppConfig) (*m.AppCreateResp_ReportApp, error) {
	now := time.Now()

	// 插入app表
	reportApp := detail.ReportApp
	app := reportApp.App
	id, err := a.AppMapper.InsertWithEcho(ctx, &appmapper.ReportApp{
		App: appmapper.AppBase{
			Name:        app.Name,
			Description: app.Description,
			Lang:        app.Lang,
			Platform:    app.Platform,
			Url:         app.Url,
			Appid:       app.Appid,
			Auth:        app.Auth,
			Stream:      app.Stream,
			Level:       app.Level,
			Status:      app.Status,
			ExpireTime:  app.ExpireTime,
			CreateTime:  now,
			UpdateTime:  now,
			Type:        consts.ReportApp,
		},
	})
	if err != nil {
		return nil, err
	}

	// 插入unit_model表
	err = a.ModelMapper.UpdateAppid(ctx, model.Id, consts.ReportApp, id)
	if err != nil {
		return nil, err
	}

	app.Id = id
	app.CreateTime = now.Unix()
	app.UpdateTime = now.Unix()
	return &m.AppCreateResp_ReportApp{
		ReportApp: &m.ReportApp{
			App: app,
		},
	}, nil
}
