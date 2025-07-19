package unit

import (
	"github.com/xh-polaris/psych-model/biz/infrastructure/config"
	"github.com/zeromicro/go-zero/core/stores/monc"
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
