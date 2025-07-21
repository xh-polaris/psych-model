package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	m "github.com/xh-polaris/psych-idl/kitex_gen/model"
	"github.com/xh-polaris/psych-model/biz/infrastructure/consts"
	untmapper "github.com/xh-polaris/psych-model/biz/infrastructure/mapper/model"
	"github.com/xh-polaris/psych-model/biz/infrastructure/util/result"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// 判断权限
	if !req.Admin {
		return nil, consts.ErrAuth
	}

	// 插入数据
	config := req.GetUnitAppConfig()
	id, err := u.UnitAppConfigMapper.InsertWithEcho(ctx, &untmapper.UnitAppConfig{
		UnitId: config.UnitId,
		Option: config.Option,
		Status: consts.Active,
	})
	if err != nil {
		return nil, err
	}

	return &m.UnitAppConfigCreateResp{
		UnitAppConfig: &m.UnitAppConfig{
			Id:     id,
			UnitId: config.UnitId,
			Option: config.Option,
			Status: consts.Active,
		},
	}, nil
}

func (u UnitAppConfigService) UnitAppConfigUpdate(ctx context.Context, req *m.UnitAppConfigUpdateReq) (res *basic.Response, err error) {
	// 判断权限
	if !req.Admin {
		return nil, consts.ErrAuth
	}

	// 更新数据
	config := req.GetUnitAppConfig()
	oid, err := primitive.ObjectIDFromHex(config.Id)
	err = u.UnitAppConfigMapper.Update(ctx, &untmapper.UnitAppConfig{
		ID:     oid,
		UnitId: config.UnitId,
		Chat:   config.Chat,
		Asr:    config.Asr,
		Tts:    config.Tts,
		Report: config.Report,
		Option: config.Option,
		Status: config.Status,
	})
	if err != nil {
		return nil, err
	}
	return result.ResponseOk(), nil
}

func (u UnitAppConfigService) UnitAppConfigGet(ctx context.Context, req *m.UnitAppConfigGetReq) (res *m.UnitAppConfigGetResp, err error) {
	// 判断权限
	if !req.Admin {
		return nil, consts.ErrAuth
	}

	// 查询数据
	appConfig, err := u.UnitAppConfigMapper.FindOneByUnitId(ctx, req.UnitId)
	if err != nil {
		return nil, err
	}
	return &m.UnitAppConfigGetResp{
		UnitAppConfig: &m.UnitAppConfig{
			Id:         appConfig.ID.Hex(),
			UnitId:     appConfig.UnitId,
			Chat:       appConfig.Chat,
			Asr:        appConfig.Asr,
			Tts:        appConfig.Tts,
			Report:     appConfig.Report,
			Option:     appConfig.Option,
			Status:     appConfig.Status,
			CreateTime: appConfig.CreateTime.Unix(),
			UpdateTime: appConfig.UpdateTime.Unix(),
		},
	}, nil
}
