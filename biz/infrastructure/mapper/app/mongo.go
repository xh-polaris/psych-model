package app

import (
	"context"
	"errors"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	"github.com/xh-polaris/psych-model/biz/infrastructure/config"
	"github.com/xh-polaris/psych-model/biz/infrastructure/consts"
	util "github.com/xh-polaris/psych-model/biz/infrastructure/util/page"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	var raws []AppWrapRaw
	err := m.conn.Find(ctx, &raws, bson.M{
		consts.ConfigId: configId,
	})
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
		ID:       raw.ID,
		Type:     raw.Type,
		ConfigId: raw.ConfigId,
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
