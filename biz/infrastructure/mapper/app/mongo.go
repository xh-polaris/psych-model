package app

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
	prefixUnitCacheKey = "cache:app"
	CollectionName     = "app"
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

func (m *MongoMapper) InsertWithEcho(ctx context.Context, app AppInterface) (string, error) {
	res, err := m.conn.InsertOneNoCache(ctx, app)
	if err != nil {
		return "", err
	}
	// 获取回显id
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (m *MongoMapper) Update(ctx context.Context, app AppInterface) error {
	app.GetBase().UpdateTime = time.Now()
	_, err := m.conn.UpdateByIDNoCache(ctx, app.GetBase().ID, bson.M{"$set": app})
	return err
}

func (m *MongoMapper) FindByUnitId(ctx context.Context, id string, t int32) (AppInterface, error) {
	var raw bson.M
	err := m.conn.FindOneNoCache(ctx, &raw, bson.M{
		consts.UnitId: id,
		consts.Type:   t,
	})
	if err != nil {
		if errors.Is(err, monc.ErrNotFound) {
			return nil, consts.ErrNotFound
		}
		return nil, err
	}

	bsonBytes, _ := bson.Marshal(raw)
	switch t {
	case consts.ChatApp:
		var app ChatApp
		if err := bson.Unmarshal(bsonBytes, &app); err != nil {
			return nil, err
		}
		return &app, nil
	case consts.TtsApp:
		var app TtsApp
		if err := bson.Unmarshal(bsonBytes, &app); err != nil {
			return nil, err
		}
		return &app, nil
	case consts.AsrApp:
		var app AsrApp
		if err := bson.Unmarshal(bsonBytes, &app); err != nil {
			return nil, err
		}
		return &app, nil
	case consts.ReportApp:
		var app ReportApp
		if err := bson.Unmarshal(bsonBytes, &app); err != nil {
			return nil, err
		}
		return &app, nil
	default:
		return nil, consts.ErrInvalidParams
	}
}

func (m *MongoMapper) Delete(ctx context.Context, unitId string, t int32) error {
	_, err := m.conn.DeleteOneNoCache(ctx, bson.M{consts.UnitId: unitId, consts.Type: t})
	return err
}
