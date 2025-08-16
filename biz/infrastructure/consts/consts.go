package consts

// 数据库相关
const (
	ID         = "_id"
	UnitId     = "unit_id"
	Status     = "status"
	CreateTime = "create_time"
	UpdateTime = "update_time"
	DeleteTime = "delete_time"
	Form       = "form"
	NotEqual   = "$ne"
	Chat       = "chat"
	Tts        = "tts"
	Asr        = "asr"
	Report     = "report"
	Type       = "type"
	ConfigId   = "config_id"
)

// status
const (
	Active  = 0
	Deleted = 1
)

// app type
const (
	All       = -1
	ChatApp   = 0
	TtsApp    = 1
	AsrApp    = 2
	ReportApp = 3
)
