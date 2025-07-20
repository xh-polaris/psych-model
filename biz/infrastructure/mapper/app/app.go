package app

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type App struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
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
	CreateTime  time.Time          `bson:"create_time,omitempty" json:"createTime,omitempty"`
	ExpireTime  time.Time          `bson:"expire_time,omitempty" json:"expireTime,omitempty"`
	UpdateTime  time.Time          `bson:"update_time,omitempty" json:"updateTime,omitempty"`
	DeleteTime  time.Time          `bson:"delete_time,omitempty" json:"deleteTime,omitempty"`
}

type ChatApp struct {
	App App `bson:"app,omitempty" json:"app,omitempty"`
}

type audioParam struct {
	Format       string `bson:"format,omitempty" json:"format,omitempty"`
	Rate         int32  `bson:"rate,omitempty" json:"rate,omitempty"`
	Bit          int32  `bson:"bit,omitempty" json:"bit,omitempty"`
	SpeechRate   int32  `bson:"speech_rate,omitempty" json:"speechRate,omitempty"`
	LoudnessRate int32  `bson:"loudness_rate,omitempty" json:"loudnessRate,omitempty"`
	Lang         string `bson:"lang,omitempty" json:"lang,omitempty"`
}
type TtsApp struct {
	App        App        `bson:"app,omitempty" json:"app,omitempty"`
	Namespace  string     `bson:"name,omitempty" json:"name,omitempty"`
	Speaker    string     `bson:"speaker,omitempty" json:"speaker,omitempty"`
	ResourceId string     `bson:"resource_id,omitempty" json:"resourceId,omitempty"`
	AudioParam audioParam `bson:"audio_param,omitempty" json:"audioParam,omitempty"`
}

type AsrApp struct {
	App        App    `bson:"app,omitempty" json:"app,omitempty"`
	Format     string `bson:"format,omitempty" json:"format,omitempty"`
	Codec      string `bson:"codec,omitempty" json:"codec,omitempty"`
	Rate       int32  `bson:"rate,omitempty" json:"rate,omitempty"`
	Bit        int32  `bson:"bit,omitempty" json:"bit,omitempty"`
	Channels   int32  `bson:"channels,omitempty" json:"channels,omitempty"`
	ModelName  string `bson:"model_name,omitempty" json:"modelName,omitempty"`
	EnablePunc bool   `bson:"enable_punc,omitempty" json:"enablePunc,omitempty"`
	EnableDdc  bool   `bson:"enable_ddc,omitempty" json:"enableDdc,omitempty"`
	ResultType string `bson:"result_type,omitempty" json:"resultType,omitempty"`
}

type RepostApp struct {
	App App `bson:"app,omitempty" json:"app,omitempty"`
}
