package model

import (
	"context"
	"errors"
	"github.com/xh-polaris/psych-model/biz/infrastructure/config"
	"github.com/xh-polaris/psych-model/biz/infrastructure/consts"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	prefixUnitCacheKey = "cache:unit_app_config"
	CollectionName     = "unit_app_config"
)

type IMongoMapper interface {
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{
		conn: conn,
	}
}

func (m *MongoMapper) InsertWithEcho(ctx context.Context, appConfig *UnitAppConfig) (string, error) {
	appConfig.ID = primitive.NewObjectID()
	appConfig.CreateTime = time.Now()
	appConfig.UpdateTime = appConfig.CreateTime
	res, err := m.conn.InsertOneNoCache(ctx, appConfig)
	if err != nil {
		return "", err
	}
	// 获取回显id
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (m *MongoMapper) Update(ctx context.Context, appConfig *UnitAppConfig) error {
	appConfig.UpdateTime = time.Now()
	_, err := m.conn.UpdateByIDNoCache(ctx, appConfig.ID, bson.M{"$set": appConfig})
	return err
}

func (m *MongoMapper) UpdateAppid(ctx context.Context, id string, types int32, appid string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	var set bson.M
	switch types {
	case consts.ChatApp:
		set = bson.M{consts.Chat: appid}
	case consts.TtsApp:
		set = bson.M{consts.Tts: appid}
	case consts.AsrApp:
		set = bson.M{consts.Asr: appid}
	case consts.ReportApp:
		set = bson.M{consts.Report: appid}
	}
	set[consts.UpdateTime] = time.Now()
	_, err = m.conn.UpdateByIDNoCache(ctx, oid, bson.M{"$set": set})
	return err
}

func (m *MongoMapper) FindOneByUnitId(ctx context.Context, id string) (*UnitAppConfig, error) {
	var appConfig UnitAppConfig
	err := m.conn.FindOneNoCache(ctx, &appConfig, bson.M{
		consts.UnitId: id,
	})
	if err != nil {
		if errors.Is(err, monc.ErrNotFound) {
			return nil, consts.ErrNotFound
		}
		return nil, err
	}
	return &appConfig, nil
}
