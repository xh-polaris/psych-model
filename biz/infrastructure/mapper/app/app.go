package app

import (
	"github.com/xh-polaris/psych-model/biz/infrastructure/consts"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AppInterface interface {
	GetBase() *AppBase
	GetType() int32
}

type AppBase struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UnitId      string             `bson:"unit_id,omitempty" json:"unitId,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Lang        string             `bson:"lang,omitempty" json:"lang,omitempty"`
	Platform    string             `bson:"platform,omitempty" json:"platform,omitempty"`
	Url         string             `bson:"url,omitempty" json:"url,omitempty"`
	Appid       string             `bson:"appid,omitempty" json:"appid,omitempty"`
	Auth        string             `bson:"auth,omitempty" json:"auth,omitempty"`
	Stream      bool               `bson:"stream,omitempty" json:"stream,omitempty"`
	Level       int32              `bson:"level,omitempty" json:"level,omitempty"`
	Status      int32              `bson:"status,omitempty" json:"status,omitempty"`
	ExpireTime  int64              `bson:"expire_time,omitempty" json:"expireTime,omitempty"`
	CreateTime  time.Time          `bson:"create_time,omitempty" json:"createTime,omitempty"`
	UpdateTime  time.Time          `bson:"update_time,omitempty" json:"updateTime,omitempty"`
	DeleteTime  time.Time          `bson:"delete_time,omitempty" json:"deleteTime,omitempty"`
	Type        int32              `bson:"type,omitempty" json:"type,omitempty"`
}

type ChatApp struct {
	App AppBase `bson:"app,omitempty" json:"app,omitempty"`
}

func (a *ChatApp) GetBase() *AppBase { return &a.App }
func (a *ChatApp) GetType() int32    { return consts.ChatApp }

type AudioParam struct {
	Format       string `bson:"format,omitempty" json:"format,omitempty"`
	Rate         int32  `bson:"rate,omitempty" json:"rate,omitempty"`
	Bit          int32  `bson:"bit,omitempty" json:"bit,omitempty"`
	SpeechRate   int32  `bson:"speech_rate,omitempty" json:"speechRate,omitempty"`
	LoudnessRate int32  `bson:"loudness_rate,omitempty" json:"loudnessRate,omitempty"`
	Lang         string `bson:"lang,omitempty" json:"lang,omitempty"`
}
type TtsApp struct {
	App        AppBase    `bson:"app,omitempty" json:"app,omitempty"`
	Namespace  string     `bson:"name,omitempty" json:"name,omitempty"`
	Speaker    string     `bson:"speaker,omitempty" json:"speaker,omitempty"`
	ResourceId string     `bson:"resource_id,omitempty" json:"resourceId,omitempty"`
	AudioParam AudioParam `bson:"audio_param,omitempty" json:"audioParam,omitempty"`
}

func (a *TtsApp) GetBase() *AppBase { return &a.App }
func (a *TtsApp) GetType() int32    { return consts.TtsApp }

type AsrApp struct {
	App        AppBase `bson:"app,omitempty" json:"app,omitempty"`
	Format     string  `bson:"format,omitempty" json:"format,omitempty"`
	Codec      string  `bson:"codec,omitempty" json:"codec,omitempty"`
	Rate       int32   `bson:"rate,omitempty" json:"rate,omitempty"`
	Bits       int32   `bson:"bits,omitempty" json:"bits,omitempty"`
	Channels   int32   `bson:"channels,omitempty" json:"channels,omitempty"`
	ModelName  string  `bson:"model_name,omitempty" json:"modelName,omitempty"`
	EnablePunc bool    `bson:"enable_punc,omitempty" json:"enablePunc,omitempty"`
	EnableDdc  bool    `bson:"enable_ddc,omitempty" json:"enableDdc,omitempty"`
	ResultType string  `bson:"result_type,omitempty" json:"resultType,omitempty"`
}

func (a *AsrApp) GetBase() *AppBase { return &a.App }
func (a *AsrApp) GetType() int32    { return consts.AsrApp }

type ReportApp struct {
	App AppBase `bson:"app,omitempty" json:"app,omitempty"`
}

func (a *ReportApp) GetBase() *AppBase { return &a.App }
func (a *ReportApp) GetType() int32    { return consts.ReportApp }
