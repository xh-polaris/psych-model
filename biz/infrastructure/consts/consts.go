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
)

// status
const (
	Active  = 0
	Deleted = 1
)

// app type
const (
	All    = 0
	Chat   = 1
	Tts    = 2
	Asr    = 3
	Repost = 4
)
