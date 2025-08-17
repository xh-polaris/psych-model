package app

import (
	"context"
	"errors"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	"github.com/xh-polaris/psych-model/biz/infrastructure/config"
	"github.com/xh-polaris/psych-model/biz/infrastructure/consts"
	"github.com/xh-polaris/psych-model/biz/infrastructure/mapper/model"
	util "github.com/xh-polaris/psych-model/biz/infrastructure/util/page"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	prefixUnitCacheKey  = "cache:app"
	AppCollectionName   = consts.AppDB
	ModelCollectionName = consts.ConfigDB
)

type IMongoMapper interface {
}

type MongoMapper struct {
	conn     *monc.Model
	linkConn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, AppCollectionName, config.Cache)
	linkConn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, ModelCollectionName, config.Cache)
	return &MongoMapper{
		conn:     conn,
		linkConn: linkConn,
	}
}

func (m *MongoMapper) InsertWithEcho(ctx context.Context, app *AppWrap) (string, error) {
	base := app.App.GetBase()
	base.CreateTime = time.Now()
	base.UpdateTime = base.CreateTime

	result, err := m.conn.InsertOneNoCache(ctx, app)
	if err != nil {
		return "", err
	}
	// 获取回显id
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (m *MongoMapper) Update(ctx context.Context, app *AppWrap) error {
	app.App.GetBase().UpdateTime = time.Now()
	_, err := m.conn.UpdateByIDNoCache(ctx, app.ID, bson.M{"$set": app})
	return err
}

func (m *MongoMapper) FindOneById(ctx context.Context, id string) (*AppWrap, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var raw AppWrapRaw
	err = m.conn.FindOneNoCache(ctx, &raw, bson.M{consts.ID: oid})
	if err != nil {
		if errors.Is(err, monc.ErrNotFound) {
			return nil, consts.ErrNotFound
		}
		return nil, err
	}

	res, err := unmarshalAppByType(&raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MongoMapper) List(ctx context.Context, p *basic.PaginationOptions, types int32) (apps []AppWrap, total int64, err error) {
	skip, limit := util.ParsePageOpt(p)
	raws := make([]AppWrapRaw, 0, limit)
	if types == consts.All {
		err = m.conn.Find(ctx, &raws,
			bson.M{}, &options.FindOptions{
				Skip:  &skip,
				Limit: &limit,
				Sort:  bson.M{consts.CreateTime: -1},
			})
		if err != nil {
			return nil, 0, err
		}
		total, err = m.conn.CountDocuments(ctx, bson.M{})
	} else {
		err = m.conn.Find(ctx, &raws,
			bson.M{
				consts.Type: types,
			}, &options.FindOptions{
				Skip:  &skip,
				Limit: &limit,
				Sort:  bson.M{consts.CreateTime: -1},
			})
		if err != nil {
			return nil, 0, err
		}
		total, err = m.conn.CountDocuments(ctx, bson.M{
			consts.Type: types,
		})
	}

	if err != nil {
		return nil, 0, err
	}

	for _, raw := range raws {
		wrap, err := unmarshalAppByType(&raw)
		if err != nil {
			return nil, 0, err
		}
		apps = append(apps, *wrap)
	}
	return apps, total, nil
}

func (m *MongoMapper) DeleteOneById(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	_, err = m.conn.DeleteOneNoCache(ctx, bson.M{consts.ID: oid})
	return err
}

func (m *MongoMapper) FindBatchByConfigId(ctx context.Context, configId string) ([]*AppWrap, error) {
	var appConfig model.UnitAppConfig
	oid, err := primitive.ObjectIDFromHex(configId)
	if err != nil {
		return nil, err
	}

	if err := m.linkConn.FindOneNoCache(ctx, &appConfig, bson.M{consts.ID: oid}); err != nil {
		return nil, err
	}
	var ids []string
	ids = append(ids, appConfig.Chat)
	ids = append(ids, appConfig.Tts)
	ids = append(ids, appConfig.Asr)
	ids = append(ids, appConfig.Report)
	var objectIDs []primitive.ObjectID
	for _, idStr := range ids {
		oid, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			continue
		}
		objectIDs = append(objectIDs, oid)
	}

	var raws []AppWrapRaw
	err = m.conn.Find(ctx, &raws, bson.M{consts.ID: bson.M{"$in": objectIDs}})
	if err != nil {
		if errors.Is(err, monc.ErrNotFound) {
			return nil, consts.ErrNotFound
		}
		return nil, err
	}

	apps := make([]*AppWrap, len(raws))
	for idx, raw := range raws {
		wrap, err := unmarshalAppByType(&raw)
		if err != nil {
			return nil, err
		}
		apps[idx] = wrap
	}

	return apps, nil
}

func unmarshalAppByType(raw *AppWrapRaw) (*AppWrap, error) {
	res := &AppWrap{
		ID:   raw.ID,
		Type: raw.Type,
	}

	switch raw.Type {
	case consts.ChatApp:
		var app ChatApp
		if err := bson.Unmarshal(raw.App, &app); err != nil {
			return nil, err
		}
		res.App = &app
	case consts.TtsApp:
		var app TtsApp
		if err := bson.Unmarshal(raw.App, &app); err != nil {
			return nil, err
		}
		res.App = &app
	case consts.AsrApp:
		var app AsrApp
		if err := bson.Unmarshal(raw.App, &app); err != nil {
			return nil, err
		}
		res.App = &app
	case consts.ReportApp:
		var app ReportApp
		if err := bson.Unmarshal(raw.App, &app); err != nil {
			return nil, err
		}
		res.App = &app
	}
	return res, nil
}
