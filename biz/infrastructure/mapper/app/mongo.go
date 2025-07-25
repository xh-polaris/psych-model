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

func (m *MongoMapper) FindByUnitId(ctx context.Context, unitId string) ([]AppInterface, error) {
	var apps []AppInterface
	err := m.conn.Find(ctx, apps, bson.M{
		consts.UnitId: unitId,
	})
	if err != nil {
		if errors.Is(err, monc.ErrNotFound) {
			return nil, consts.ErrNotFound
		}
		return nil, err
	}

	return apps, nil
}

func (m *MongoMapper) FindPagination(ctx context.Context, p *basic.PaginationOptions, types int32) (apps []AppInterface, total int64, err error) {
	skip, limit := util.ParsePageOpt(p)
	apps = make([]AppInterface, 0, limit)
	if types == consts.All {
		err = m.conn.Find(ctx, &apps,
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
		err = m.conn.Find(ctx, &apps,
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

	return apps, total, nil
}

func (m *MongoMapper) Delete(ctx context.Context, unitId string, t int32) error {
	_, err := m.conn.DeleteOneNoCache(ctx, bson.M{consts.UnitId: unitId, consts.Type: t})
	return err
}
