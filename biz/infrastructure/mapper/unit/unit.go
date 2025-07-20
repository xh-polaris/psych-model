package unit

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UnitAppConfig struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UnitId     string             `bson:"unit_id,omitempty" json:"unit_id,omitempty"`
	Chat       string             `bson:"chat,omitempty" json:"chat,omitempty"`
	Asr        string             `bson:"asr,omitempty" json:"asr,omitempty"`
	Tts        string             `bson:"tts,omitempty" json:"tts,omitempty"`
	Report     string             `bson:"report,omitempty" json:"report,omitempty"`
	Option     map[string]string  `bson:"option,omitempty" json:"option,omitempty"`
	Status     int32              `bson:"status,omitempty" json:"status,omitempty"`
	CreateTime time.Time          `bson:"create_time,omitempty" json:"createTime,omitempty"`
	UpdateTime time.Time          `bson:"update_time,omitempty" json:"updateTime,omitempty"`
}
