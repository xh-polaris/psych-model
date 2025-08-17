package app

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AppInterface interface {
	GetBase() *AppBase
}

type AppWrap struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type int32              `bson:"type,omitempty"`
	App  AppInterface       `bson:"app,omitempty"`
}

type AppWrapRaw struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type int32              `bson:"type,omitempty"`
	App  bson.Raw           `bson:"app,omitempty"`
}

type AppBase struct {
	Name        string    `bson:"name,omitempty" json:"name,omitempty"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	Lang        string    `bson:"lang,omitempty" json:"lang,omitempty"`
	Platform    string    `bson:"platform,omitempty" json:"platform,omitempty"`
	Provider    string    `bson:"provider,omitempty" json:"provider,omitempty"`
	Url         string    `bson:"url,omitempty" json:"url,omitempty"`
	AppId       string    `bson:"app_id,omitempty" json:"appId,omitempty"`
	AccessKey   string    `bson:"access_key,omitempty" json:"accessKey,omitempty"`
	Stream      bool      `bson:"stream,omitempty" json:"stream,omitempty"`
	Level       int32     `bson:"level,omitempty" json:"level,omitempty"`
	Status      int32     `bson:"status,omitempty" json:"status,omitempty"`
	CreateTime  time.Time `bson:"create_time,omitempty" json:"createTime,omitempty"`
	UpdateTime  time.Time `bson:"update_time,omitempty" json:"updateTime,omitempty"`
	DeleteTime  time.Time `bson:"delete_time,omitempty" json:"deleteTime,omitempty"`
}

type ChatApp struct {
	AppBase `bson:",inline"`
}

type AudioParam struct {
	Format       string `bson:"format,omitempty" json:"format,omitempty"`
	Codec        string `bson:"codec,omitempty" json:"codec,omitempty"`
	Rate         int32  `bson:"rate,omitempty" json:"rate,omitempty"`
	Bits         int32  `bson:"bits,omitempty" json:"bits,omitempty"`
	Channels     int32  `bson:"channels,omitempty" json:"channels,omitempty"`
	SpeechRate   int32  `bson:"speech_rate,omitempty" json:"speechRate,omitempty"`
	LoudnessRate int32  `bson:"loudness_rate,omitempty" json:"loudnessRate,omitempty"`
	Lang         string `bson:"lang,omitempty" json:"lang,omitempty"`
	ResultType   string `bson:"result_type,omitempty" json:"resultType,omitempty"`
}
type TtsApp struct {
	AppBase    `bson:",inline"`
	Namespace  string     `bson:"name,omitempty" json:"name,omitempty"`
	Speaker    string     `bson:"speaker,omitempty" json:"speaker,omitempty"`
	ResourceId string     `bson:"resource_id,omitempty" json:"resourceId,omitempty"`
	AudioParam AudioParam `bson:"audio_param,omitempty" json:"audioParam,omitempty"`
}

type AsrApp struct {
	AppBase    `bson:",inline"`
	Format     string `bson:"format,omitempty" json:"format,omitempty"`
	Codec      string `bson:"codec,omitempty" json:"codec,omitempty"`
	Rate       int32  `bson:"rate,omitempty" json:"rate,omitempty"`
	Bits       int32  `bson:"bits,omitempty" json:"bits,omitempty"`
	ResourceId string `bson:"resource_id,omitempty" json:"resourceId,omitempty"`
	Channels   int32  `bson:"channels,omitempty" json:"channels,omitempty"`
	ModelName  string `bson:"model_name,omitempty" json:"modelName,omitempty"`
	EnablePunc bool   `bson:"enable_punc,omitempty" json:"enablePunc,omitempty"`
	EnableDdc  bool   `bson:"enable_ddc,omitempty" json:"enableDdc,omitempty"`
	ResultType string `bson:"result_type,omitempty" json:"resultType,omitempty"`
}

type ReportApp struct {
	AppBase `bson:",inline"`
}

func (a *ChatApp) GetBase() *AppBase   { return &a.AppBase }
func (a *TtsApp) GetBase() *AppBase    { return &a.AppBase }
func (a *AsrApp) GetBase() *AppBase    { return &a.AppBase }
func (a *ReportApp) GetBase() *AppBase { return &a.AppBase }
