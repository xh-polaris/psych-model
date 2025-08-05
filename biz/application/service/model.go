package service

import (
	"context"
	"errors"
	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	m "github.com/xh-polaris/psych-idl/kitex_gen/model"
	"github.com/xh-polaris/psych-model/biz/infrastructure/consts"
	appmapper "github.com/xh-polaris/psych-model/biz/infrastructure/mapper/app"
	mdlmapper "github.com/xh-polaris/psych-model/biz/infrastructure/mapper/model"
	"github.com/xh-polaris/psych-model/biz/infrastructure/util/result"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUnitAppConfigService interface {
	UnitAppConfigCreate(ctx context.Context, req *m.UnitAppConfigCreateReq) (res *m.UnitAppConfigCreateResp, err error)
	UnitAppConfigUpdate(ctx context.Context, req *m.UnitAppConfigUpdateReq) (res *basic.Response, err error)
	UnitAppConfigGetByUnitId(ctx context.Context, req *m.UnitAppConfigGetByUnitIdReq) (res *m.UnitAppConfigGetByUnitIdResp, err error)
	UnitAppConfigGetById(ctx context.Context, req *m.UnitAppConfigGetByIdReq) (res *m.UnitAppConfigGetByIdResp, err error)
}

type UnitAppConfigService struct {
	ModelMapper *mdlmapper.MongoMapper
	AppMapper   *appmapper.MongoMapper
}

var UnitAppConfigServiceSet = wire.NewSet(
	wire.Struct(new(UnitAppConfigService), "*"),
	wire.Bind(new(IUnitAppConfigService), new(*UnitAppConfigService)),
)

func (u *UnitAppConfigService) UnitAppConfigCreate(ctx context.Context, req *m.UnitAppConfigCreateReq) (res *m.UnitAppConfigCreateResp, err error) {
	// 判断权限
	if !req.Admin {
		return nil, consts.ErrAuth
	}

	// 查看是否已经存在unitId重复的配置
	config := req.GetUnitAppConfig()
	unitId := config.UnitId
	if _, err = u.ModelMapper.FindOneByUnitId(ctx, unitId); err == nil || !errors.Is(err, consts.ErrNotFound) {
		// err == nil 无报错 -> 有数据
		// err != nil, 且不是 NotFound, 发生了其他错误
		return nil, consts.ErrExistConfig
	}

	// 插入数据
	id, err := u.ModelMapper.InsertWithEcho(ctx, &mdlmapper.UnitAppConfig{
		UnitId: config.UnitId,
		Name:   config.Name,
		Video:  config.Video,
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
			Name:   config.Name,
			Video:  config.Video,
			Option: config.Option,
			Status: consts.Active,
		},
	}, nil
}

func (u *UnitAppConfigService) UnitAppConfigUpdate(ctx context.Context, req *m.UnitAppConfigUpdateReq) (res *basic.Response, err error) {
	// 判断权限
	if !req.Admin {
		return nil, consts.ErrAuth
	}

	// 更新数据
	config := req.GetUnitAppConfig()
	oid, err := primitive.ObjectIDFromHex(config.Id)
	err = u.ModelMapper.Update(ctx, &mdlmapper.UnitAppConfig{
		ID:     oid,
		Name:   config.Name,
		Video:  config.Video,
		Option: config.Option,
	})
	if err != nil {
		return nil, err
	}
	return result.ResponseOk(), nil
}

func (u *UnitAppConfigService) UnitAppConfigGetByUnitId(ctx context.Context, req *m.UnitAppConfigGetByUnitIdReq) (res *m.UnitAppConfigGetByUnitIdResp, err error) {
	// 查询model
	appConfig, err := u.ModelMapper.FindOneByUnitId(ctx, req.UnitId)
	if err != nil {
		return nil, err
	}
	// 查询app
	apps, err := u.AppMapper.FindBatchByConfigId(ctx, appConfig.ID.Hex())
	if err != nil && !errors.Is(err, consts.ErrNotFound) {
		return nil, err
	}

	resApps := make([]*m.AppData, len(apps))
	for idx, app := range apps {
		gen := appDB2Gen(app, req.Admin)
		resApps[idx] = gen
	}
	return &m.UnitAppConfigGetByUnitIdResp{
		UnitAppConfig: configDB2Gen(appConfig),
		Apps:          resApps,
	}, nil
}

func (u *UnitAppConfigService) UnitAppConfigGetById(ctx context.Context, req *m.UnitAppConfigGetByIdReq) (res *m.UnitAppConfigGetByIdResp, err error) {
	appConfig, err := u.ModelMapper.FindOneById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	apps, err := u.AppMapper.FindBatchByConfigId(ctx, req.Id)
	if err != nil {
		return &m.UnitAppConfigGetByIdResp{
			UnitAppConfig: configDB2Gen(appConfig),
		}, nil
	}

	resApps := make([]*m.AppData, len(apps))
	for idx, app := range apps {
		gen := appDB2Gen(app, req.Admin)
		resApps[idx] = gen
	}

	return &m.UnitAppConfigGetByIdResp{
		UnitAppConfig: configDB2Gen(appConfig),
		Apps:          resApps,
	}, nil
}

func configDB2Gen(db *mdlmapper.UnitAppConfig) *m.UnitAppConfig {
	return &m.UnitAppConfig{
		Id:         db.ID.Hex(),
		UnitId:     db.UnitId,
		Name:       db.Name,
		Video:      db.Video,
		Chat:       db.Chat,
		Asr:        db.Asr,
		Tts:        db.Tts,
		Report:     db.Report,
		Option:     db.Option,
		Status:     db.Status,
		CreateTime: db.CreateTime.Unix(),
		UpdateTime: db.UpdateTime.Unix(),
	}
}
